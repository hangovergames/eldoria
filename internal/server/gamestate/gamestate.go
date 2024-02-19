// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gamestate

import (
	"fmt"
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

func (m *GameState) FindPlayer(name string) (game.IPlayer, error) {
	for _, player := range m.Players {
		if player.GetName() == name {
			return player, nil
		}
	}
	return nil, fmt.Errorf("player with name %s not found", name)
}
