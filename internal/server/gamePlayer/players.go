// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gamePlayer

type PlayerID uint

type Player struct {
	ID        PlayerID       // ID Unique identifier for the gamePlayer.
	Name      string         // Player's name or username.
	Resources map[string]int // Resources owned by the gamePlayer, keyed by resource name.
}

// NewPlayer creates a new Player instance with initial values.
func NewPlayer(id PlayerID, name string) *Player {
	return &Player{
		ID:        id,
		Name:      name,
		Resources: make(map[string]int), // Initialize with no resources.
	}
}
