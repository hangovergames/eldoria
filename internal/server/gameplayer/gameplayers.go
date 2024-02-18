// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameplayer

import "github.com/hangovergames/eldoria/internal/server/game"

type Player struct {
	ID        game.PlayerID  // ID Unique identifier for the gamePlayer.
	Name      string         // Player's name or username.
	Resources map[string]int // Resources owned by the gamePlayer, keyed by resource name.
}

// NewPlayer creates a new Player instance with initial values.
func NewPlayer(id game.PlayerID, name string) *Player {
	return &Player{
		ID:        id,
		Name:      name,
		Resources: make(map[string]int), // Initialize with no resources.
	}
}

func (p *Player) GetID() game.PlayerID {
	return p.ID
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) SetName(name string) {
	p.Name = name
}

func (p *Player) GetResources() map[string]int {
	return p.Resources
}

func (p *Player) AddResource(resource string, amount int) {
	p.Resources[resource] += amount
}

func (p *Player) SubtractResource(resource string, amount int) {
	p.Resources[resource] -= amount
	if p.Resources[resource] < 0 {
		p.Resources[resource] = 0
	}
}
