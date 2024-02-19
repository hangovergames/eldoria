// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package players

import (
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/hangovergames/eldoria/internal/server/game"
	"net/http"
)

// PlayerCheck handles the GET /players/:username which is used to check player
// availability before registration or sign in attempt
func PlayerCheck(
	response game.IResponse,
	request game.IRequest,
	server game.IServer,
) {

	if !request.IsMethodGet() {
		response.SendMethodNotSupportedError()
		return
	}

	params := request.GetVars()
	username := params["username"]
	state := server.GetState()

	player, err := state.FindPlayer(username)

	if player == nil || err != nil {
		response.SendError(404, "Not Found")
		return
	}

	data := dtos.NewPlayerDTO(player.GetName(), 0, 0)

	response.Send(http.StatusOK, data)

}
