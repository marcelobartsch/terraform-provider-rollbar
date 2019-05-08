package rollbar

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ListProjectsResponse represents the list projects response
type ListProjectsResponse struct {
	Error  int `json:"err"`
	Result []struct {
		AccountID    int    `json:"account_id"`
		ID           int    `json:"id"`
		Email        string `json:"email"`
		DateCreated  int    `json:"date_created"`
		DateModified int    `json:"date_modified"`
		Name         string `json:"name"`
	}
}

// ListProjects lists the projects for this API Key
func (c *Client) ListProjects() (*ListProjectsResponse, error) {
	var data ListProjectsResponse

	url := fmt.Sprintf("%sprojects?access_token=%s", c.APIBaseURL, c.APIKey)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	bytes, err := c.makeRequest(req)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
