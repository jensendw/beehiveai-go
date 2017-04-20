package bhive

import (
	"fmt"
)

// Integration JSON data that is POST'ed to API
type Integration struct {
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
	Tags  string `json:"tags,omitempty"`
}

// IntegrationID JSON data is sent to create integration id
type IntegrationID struct {
	IntegrationID string `json:"integration_id,omitempty"`
}

// IntegrationResponse is the expected response from the api
type IntegrationResponse struct {
	Message string `json:"message"`
	OK      bool   `json:"ok"`
}

// CreateIntegration adds the data for the integration
func (c *Client) CreateIntegration(title string, text string, tags string) (*IntegrationResponse, error) {
	t := Integration{Title: title, Text: text, Tags: tags}
	resp, err := c.put("/"+t.Title, t, nil)
	if err != nil {
		fmt.Println("Error decoding JSON response from creating integration: ", err)
		return nil, err
	}

	var integrationResponse IntegrationResponse

	err = c.decodeJSON(resp, &integrationResponse)
	if err != nil {
		fmt.Println("Error decoding JSON response from creating integration: ", err)
		return nil, err
	}

	integrationExists := doesIntegrationExist(&integrationResponse)
	if integrationExists == false {
		i := IntegrationID{IntegrationID: t.Title}
		_, err := c.post("/create", i)
		if err != nil {
			fmt.Println("Problem creating integration ID")
			return nil, err
		}
		resp2, err := c.put("/"+t.Title, t, nil)
		if err != nil {
			fmt.Println("Got error while creating integration after ID was created: ", err)
			return nil, err
		}

		err = c.decodeJSON(resp2, &integrationResponse)
		if err != nil {
			fmt.Println("Got error while decoding JSON response after creating ID and integration: ", err)
			return nil, err
		}
		return &integrationResponse, nil
	}
	return &integrationResponse, nil
}

// doesIntegrationExist verifies if the error response from the API
// translates to not existing
func doesIntegrationExist(intergrationResponse *IntegrationResponse) bool {
	if intergrationResponse.OK == true {
		return true
	}
	return false
}

// CreateIntegrationID calls the API to create integration id
func (c *Client) CreateIntegrationID(integrationID string) (*IntegrationResponse, error) {
	t := IntegrationID{IntegrationID: integrationID}
	resp, err := c.post("/create", t)
	if err != nil {
		fmt.Println("got error when creating integration id: ", err)
	}

	var integrationResponse IntegrationResponse

	err = c.decodeJSON(resp, &integrationResponse)
	if err != nil {
		fmt.Println("Got error while decoding JSON response after creating ID and integration: ", err)
		return nil, err
	}
	return &integrationResponse, nil
}
