// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameMap

// TileType represents different types of tiles on the game map.
type TileType int

const (
	UnknownTileType TileType = iota // Special tile, not known to the player
	DeepOcean                       // Great for impassable or difficult-to-navigate areas.
	ShallowOcean                    // Could be used for areas that allow for naval passage but perhaps not settlement.
	Lake                            // Offers fresh water, possibly influencing settlement locations or resource availability.
	Desert                          // Challenging terrain with limited resources, affecting movement and settlement.
	Swamp                           // Difficult terrain for movement, could have unique resources or units.
	Plains                          // Cold, barren lands found in northern regions. Could affect unit movement and settlement negatively but might have unique strategic or resource values.
	Grassland                       // Ideal for agriculture and settlement, probably the most basic terrain for expansion.
	Steppe                          // Vast, open grasslands with dry conditions, ideal for nomadic lifestyles and fast unit movement, but may offer limited agricultural resources.
	Hills                           // Offers defensive bonuses and possibly mineral resources.
	Mountain                        // Impassable or challenging to navigate, but could be rich in resources.
	Tundra                          // Cold, barren lands found in northern regions. Could affect unit movement and settlement negatively but might have unique strategic or resource values.
	Ice                             // Similar to mountains in being impassable, representing glaciers or frozen seas that could be navigated only by specific units or under certain conditions.
	Forest                          // Provides wood resources, potentially slows down movement.
	Savannah                        // Similar to grasslands but with different strategic or resource implications, possibly found in warmer climates.
	Volcano                         // A unique terrain type that could offer rich mineral resources but comes with risks, such as periodic eruptions that could damage nearby units or settlements.
	Jungle                          // Dense and rich in biodiversity, jungles can slow down movement due to thick vegetation but might offer unique resources and hideouts.
	Beach                           // Transitional zones between land and water, beaches can facilitate naval landings and tourism, affecting economy and military strategies.
)

// ModifierType represents different modifiers of tiles on the game map.
type ModifierType int

const (
	River      ModifierType = iota // Rivers
	Road                           // Roads
	CoralReefs                     // Coral reefs
)

// Tile represents a tile on the game map, including its type and any modifiers.
type Tile struct {
	Type      TileType
	Modifiers []ModifierType // Slice of modifiers affecting this tile
}

// NewTile creates a new Tile with the given type and modifiers.
// It returns a Tile struct.
func NewTile(tileType TileType, modifiers ...ModifierType) Tile {
	return Tile{
		Type:      tileType,
		Modifiers: modifiers,
	}
}
