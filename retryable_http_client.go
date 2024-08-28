package statuspage

import (
	"context"
	"github.com/hashicorp/go-retryablehttp"
	"net/http"
	"strconv"
	"time"
)

func NewRetryableClient() *retryablehttp.Client {
	client := retryablehttp.NewClient()
	client.CheckRetry = retryPolicy
	client.Backoff = backoffPolicy
	return client
}

const (
	RateLimitExceededStatusCode = 420
)

// We're wrapping original retryablehttp.DefaultRetryPolicy and adding retry on 420 HTTP code
func retryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	defaultShouldRetry, _ := retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	isRateLimited := resp != nil && resp.StatusCode == RateLimitExceededStatusCode
	shouldRetry := defaultShouldRetry || isRateLimited
	return shouldRetry, nil
}

// We're wrapping original retryablehttp.DefaultBackoff and using response header in such format `Retry-After: 60`
// retryablehttp already implements this behaviour, but it's executed only for 429 and 503 HTTP codes, and we need it for 420 HTTP code as well
func backoffPolicy(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	if resp != nil {
		if resp.StatusCode == RateLimitExceededStatusCode {
			if sleep, ok := parseRetryAfterHeader(resp.Header["Retry-After"]); ok {
				return sleep
			}
		}
	}
	defaultSleep := retryablehttp.DefaultBackoff(min, max, attemptNum, resp)
	return defaultSleep
}

// Code partially copied from retryablehttp.parseRetryAfterHeader
func parseRetryAfterHeader(headers []string) (time.Duration, bool) {
	if len(headers) == 0 || headers[0] == "" {
		return 0, false
	}
	header := headers[0]
	// Retry-After: 60
	if sleep, err := strconv.ParseInt(header, 10, 64); err == nil {
		if sleep < 0 { // a negative sleep doesn't make sense
			return 0, false
		}
		return time.Second * time.Duration(sleep), true
	}
	return 0, true
}
