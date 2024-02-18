// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apiclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/stretchr/testify/assert"
)

func TestFetchUIConfigDTO(t *testing.T) {

	// Setup
	expectedConfig := dtos.UIConfigDTO{
		// Populate with expected data
	}

	// Start a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with the expected configuration JSON
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedConfig)
	}))
	defer server.Close()

	// Create an instance of your APIClient pointing to the mock server
	client := NewAPIClient(server.URL)

	// Execute
	config, err := client.FetchUIConfigDTO()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, *config)

}
