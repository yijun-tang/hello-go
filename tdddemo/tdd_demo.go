package tdddemo

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data string `json:"data"`
}

type Client struct {
	apiBasePath string
	client      *http.Client
}

func (c *Client) GetMessage(url string) (string, int, error) {
	request, err := http.NewRequest(http.MethodGet, c.apiBasePath+url, nil)
	if err != nil {
		return `{"error":"record not found"}`, http.StatusInternalServerError, err
	}

	resp, err := c.client.Do(request)
	if err != nil {
		return `{"error":"record not found"}`, http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return `{"error":"record not found"}`, http.StatusNotFound, nil
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return `{"error":"record not found"}`, http.StatusInternalServerError, err
	}

	return response.Data, http.StatusOK, nil
}
