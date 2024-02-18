// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package game

type TileType uint
type ModifierType uint
type TileEffect uint

// HasTileEffect Function to check if a tile effect includes a specific effect
func (t TileEffect) HasTileEffect(effect TileEffect) bool {
	return t&effect != 0
}

// IServer defines the methods available from the Server
// that are needed by the HTTP handlers.
type IServer interface {
	Start() error
	SetupRoutes()
	GetAddress() string
	GetRuleset() IRuleset
}

// IRuleset defines the methods that the Ruleset needs to expose to external consumers.
type IRuleset interface {
	FindTileType(tileName string) (TileType, bool)
	FindModifierType(tileName string) (ModifierType, bool)
	CreateTileFromName(tileName string) (ITile, error)
	IsEffectEnabled(effectStr string) bool
	ConvertAndCombineTileEffects(effectStrings []string) (TileEffect, error)
}

// ITile defines the interface for operations on a Tile.
type ITile interface {

	// GetType returns the type of the tile.
	GetType() TileType

	// GetModifiers returns the slice of ModifierType applied to the tile.
	GetModifiers() []ModifierType

	// GetEffects returns the TileEffect applied to the tile.
	GetEffects() TileEffect

	// Clone creates and returns a deep copy of the tile.
	Clone() ITile
}
