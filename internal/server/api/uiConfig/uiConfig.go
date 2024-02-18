// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uiConfig

import (
	"github.com/hangovergames/eldoria/internal/common/dtos" // Adjust the import path based on your project structure
	"github.com/hangovergames/eldoria/internal/server/apiRequests"
	"github.com/hangovergames/eldoria/internal/server/apiResponses"
	"net/http"
)

// UIConfig handles the requests for UI configuration data.
func UIConfig(response apiResponses.Response, request apiRequests.Request) {

	if !request.IsMethodGet() {
		response.SendMethodNotSupportedError()
		return
	}

	// For demonstration, let's load the UIConfigDTO from a predefined JSON file.
	// In a real application, this might involve fetching the configuration from a database or generating it dynamically.
	uiConfig := dtos.LoadUIConfigDTO("./examples/ui-config-dto.json")

	response.Send(http.StatusOK, uiConfig)

}
