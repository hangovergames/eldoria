// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package game

import (
	"github.com/hangovergames/eldoria/internal/server/gameMap"
	"github.com/hangovergames/eldoria/internal/server/gamePlayer"
)

type Game struct {
	Map     gameMap.GameMap
	Players []gamePlayer.Player
}

func NewGameMap(
	gameMap gameMap.GameMap,
) *Game {
	return &Game{
		Map:     gameMap,
		Players: []gamePlayer.Player{},
	}
}
