// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package index

import (
	"github.com/hangovergames/eldoria/internal/server/apiRequests"
	"github.com/hangovergames/eldoria/internal/server/apiResponses"
	"github.com/hangovergames/eldoria/internal/server/game"
	"net/http"
)

// Index handles the GET requests at the root URL.
func Index(response apiResponses.Response, request apiRequests.Request, server game.IServer) {

	if !request.IsMethodGet() {
		response.SendMethodNotSupportedError()
		return
	}

	data := newIndexDTO("0.0.1")

	response.Send(http.StatusOK, data)

}