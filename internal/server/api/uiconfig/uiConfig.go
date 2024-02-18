// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uiconfig

import (
	"github.com/hangovergames/eldoria/internal/server/apirequests"
	"github.com/hangovergames/eldoria/internal/server/apiresponses"
	"github.com/hangovergames/eldoria/internal/server/game"
	"log"
	"net/http"
)

// UIConfig handles the requests for UI configuration data.
func UIConfig(response apiresponses.Response, request apirequests.Request, server game.IServer) {
	if !request.IsMethodGet() {
		response.SendMethodNotSupportedError()
		return
	}
	ruleset := server.GetRuleset()
	log.Println("ruleset = ", ruleset)
	uiConfig := ruleset.GetUI()
	log.Println("uiConfig = ", uiConfig)
	response.Send(http.StatusOK, uiConfig)
}
