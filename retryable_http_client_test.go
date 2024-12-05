package statuspage

import (
	"net/http"
	"testing"
)

func TestBackOffPolicy(t *testing.T) {
	resp := http.Response{StatusCode: StatusRateLimitExceeded}
	resp.Header = map[string][]string{}
	resp.Header.Set("Retry-After", "30")

	d := backoffPolicy(1, 50, 2, &resp)

	if d.String() != "30s" {
		t.Errorf("%s", d.String())
	}
}
