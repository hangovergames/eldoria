// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameState

import (
	"github.com/hangovergames/eldoria/internal/server/gameMap"
	"github.com/hangovergames/eldoria/internal/server/gamePlayer"
)

type GameState struct {
	Map     gameMap.GameMap
	Players []gamePlayer.Player
}

func NewGameState(
	gameMap gameMap.GameMap,
) *GameState {
	return &GameState{
		Map:     gameMap,
		Players: []gamePlayer.Player{},
	}
}
