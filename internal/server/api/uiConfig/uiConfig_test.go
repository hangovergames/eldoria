package uiConfig

import (
	"encoding/json"
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/hangovergames/eldoria/internal/server/apiRequests"
	"github.com/hangovergames/eldoria/internal/server/apiResponses"
	"github.com/hangovergames/eldoria/internal/server/game/mocks"
	"testing"
)

func TestUIConfig(t *testing.T) {
	tests := []struct {
		name          string
		requestMethod bool // true for GET, false otherwise
		expectError   bool
	}{
		{"GET Request", true, false},
		{"Non-GET Request", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockResponse := &apiResponses.MockResponse{}
			mockRequest := &apiRequests.MockRequest{IsGet: tt.requestMethod}
			mockServer := mocks.NewMockServer()

			UIConfig(mockResponse, mockRequest, mockServer)

			if tt.expectError {
				if !mockResponse.MethodNotAllowed {
					t.Errorf("Expected method not allowed error, but didn't get one")
				}
			} else {
				if mockResponse.SentStatusCode != 200 {
					t.Errorf("Expected status code 200, got %d", mockResponse.SentStatusCode)
				}
				// Assuming the SentData was marshaled into JSON, now unmarshal for assertion
				data, _ := json.Marshal(mockResponse.SentData) // Ignoring error for brevity
				var actualConfig dtos.UIConfigDTO
				if err := json.Unmarshal(data, &actualConfig); err != nil {
					t.Fatalf("Failed to unmarshal SentData: %v", err)
				}

				// Perform your assertions based on the actualConfig
				// Example assertion (adjust as needed)
				if len(actualConfig.SpriteSheets) == 0 || len(actualConfig.TileConfigs) == 0 {
					t.Errorf("Expected non-empty configurations for sprite sheets and tile configs")
				}

				// More detailed assertions can be added here based on the expected contents of your UIConfigDTO
			}
		})
	}
}
