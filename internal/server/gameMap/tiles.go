// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameMap

const UnknownModifierType = 0
const UnknownTileType = 0

type TileType uint
type ModifierType uint

// Tile represents a tile on the game map, including its type and any modifiers.
type Tile struct {
	Type      TileType       // Type is an ID of the tile from the Ruleset
	Modifiers []ModifierType // Modifiers Enabled modifiers
	Effects   TileEffect     // Effects combined active effects for the tile
}

// NewTile creates a new Tile with the given type and modifiers.
// It returns a Tile struct.
func NewTile(
	tileType TileType,
	modifiers []ModifierType,
	effects TileEffect,
) Tile {
	return Tile{
		Type:      tileType,
		Modifiers: modifiers,
		Effects:   effects,
	}
}

func (t Tile) Clone() Tile {

	// Create a new slice for Modifiers with the same length as the original
	modifiersCopy := make([]ModifierType, len(t.Modifiers))
	copy(modifiersCopy, t.Modifiers)

	// Return a new Tile with the same Type, a copy of Modifiers, and the same Effects
	return Tile{
		Type:      t.Type,
		Modifiers: modifiersCopy,
		Effects:   t.Effects,
	}

}
