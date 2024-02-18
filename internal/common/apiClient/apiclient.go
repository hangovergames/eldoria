// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apiClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hangovergames/eldoria/internal/common/dtos"
)

// IAPIClient defines the behavior for an API client that can fetch UI configurations.
type IAPIClient interface {
	FetchUIConfigDTO() (*dtos.UIConfigDTO, error)
}

// APIClient represents a client for making API requests.
type APIClient struct {
	BaseURL string // BaseURL is the root URL for the API server.
}

// NewAPIClient creates a new instance of APIClient.
func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
	}
}

// FetchUIConfigDTO fetches the UI configuration data from the server.
func (client *APIClient) FetchUIConfigDTO() (*dtos.UIConfigDTO, error) {

	// Construct the full URL for the UI config endpoint
	url := fmt.Sprintf("%s/ui/config", client.BaseURL)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request to UI config endpoint: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal the JSON response into the UIConfigDTO struct
	var config dtos.UIConfigDTO
	if err := json.Unmarshal(body, &config); err != nil {
		return nil, fmt.Errorf("error unmarshaling response into UIConfigDTO: %v", err)
	}

	return &config, nil
}
