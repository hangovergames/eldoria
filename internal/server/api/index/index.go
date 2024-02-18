// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package index

import (
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/hangovergames/eldoria/internal/server/apirequests"
	"github.com/hangovergames/eldoria/internal/server/apiresponses"
	"github.com/hangovergames/eldoria/internal/server/game"
	"net/http"
)

// Index handles the GET requests at the root URL.
func Index(response apiresponses.Response, request apirequests.Request, server game.IServer) {

	if !request.IsMethodGet() {
		response.SendMethodNotSupportedError()
		return
	}

	data := dtos.NewIndexDTO("0.0.1")

	response.Send(http.StatusOK, data)

}
