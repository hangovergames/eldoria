// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gamestate

import (
	"github.com/hangovergames/eldoria/internal/server/game"
)

type GameState struct {
	Map     game.IGameMap
	Players []game.IPlayer
}

func NewGameState(
	gameMap game.IGameMap,
) *GameState {
	return &GameState{
		Map:     gameMap,
		Players: []game.IPlayer{},
	}
}

func (m *GameState) GetMap() game.IGameMap {
	return m.Map
}

func (m *GameState) GetPlayers() []game.IPlayer {
	return m.Players
}
