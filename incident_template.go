package statuspagesdk

type IncidentTemplate struct {
	GroupID      string `json:"group_id"`
	Name         string `json:"name"`
	UpdateStatus string `json:"update_status"`
	Title        string `json:"suffix"`
	Body         int32  `json:"y_axis_min"`
	ComponentIDs int32  `json:"component_ids"`
	ShouldTweet  int32  `json:"should_tweet"`
}

type IncidentTemplateFull struct {
	IncidentTemplate
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreateIncidentTemplate(client *Client, pageID, incidentTemplate *IncidentTemplate) (*IncidentTemplateFull, error) {
	var i IncidentTemplateFull
	err := createResource(
		client,
		pageID,
		"incident_template",
		struct {
			IncidentTemplate *IncidentTemplate `json:"incident_template"`
		}{incidentTemplate},
		&i,
	)

	return &i, err
}

func GetIncidentTemplate(client *Client, pageID, incidentTemplateID string) (*IncidentTemplateFull, error) {
	var i IncidentTemplateFull
	err := readResource(client, pageID, incidentTemplateID, "incident_template", &i)

	return &i, err
}

func UpdateIncidentTemplate(client *Client, pageID, incidentTemplateID string, incidentTemplate *IncidentTemplate) (*IncidentTemplateFull, error) {
	var i IncidentTemplateFull

	err := updateResource(
		client,
		pageID,
		"incident_template",
		incidentTemplateID,
		struct {
			IncidentTemplate *IncidentTemplate `json:"incident_template"`
		}{incidentTemplate},
		&i,
	)

	return &i, err
}

func DeleteIncidentTemplate(client *Client, pageID, incidentTemplateID string) (err error) {
	return deleteResource(client, pageID, "incident_template", incidentTemplateID)
}
