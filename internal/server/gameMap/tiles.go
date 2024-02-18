// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameMap

import "github.com/hangovergames/eldoria/internal/server/game"

// Tile represents a tile on the game map, including its type and any modifiers.
type Tile struct {
	Type      game.TileType       // Type is an ID of the tile from the Ruleset
	Modifiers []game.ModifierType // Modifiers Enabled modifiers
	Effects   game.TileEffect     // Effects combined active effects for the tile
}

// NewTile creates a new Tile with the given type and modifiers.
// It returns a Tile struct.
func NewTile(
	tileType game.TileType,
	modifiers []game.ModifierType,
	effects game.TileEffect,
) Tile {
	return Tile{
		Type:      tileType,
		Modifiers: modifiers,
		Effects:   effects,
	}
}

func (t Tile) Clone() game.ITile {

	// Create a new slice for Modifiers with the same length as the original
	modifiersCopy := make([]game.ModifierType, len(t.Modifiers))
	copy(modifiersCopy, t.Modifiers)

	// Return a new Tile with the same Type, a copy of Modifiers, and the same Effects
	return Tile{
		Type:      t.Type,
		Modifiers: modifiersCopy,
		Effects:   t.Effects,
	}

}

func (t Tile) GetType() game.TileType {
	return t.Type
}

func (t Tile) GetModifiers() []game.ModifierType {
	return t.Modifiers
}

func (t Tile) GetEffects() game.TileEffect {
	return t.Effects
}
