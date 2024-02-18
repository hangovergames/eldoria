// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package players

import (
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/hangovergames/eldoria/internal/server/apirequests"
	"github.com/hangovergames/eldoria/internal/server/apiresponses"
	"github.com/hangovergames/eldoria/internal/server/game"
	"net/http"
)

// Player handles the GET requests at the root URL.
func Player(response apiresponses.Response, request apirequests.Request, server game.IServer) {

	if !request.IsMethodGet() {
		response.SendMethodNotSupportedError()
		return
	}

	//state := server.GetState()

	data := dtos.NewPlayerDTO("name", 0, 0)

	response.Send(http.StatusOK, data)

}
