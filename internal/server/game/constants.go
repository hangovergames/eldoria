// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package game

const UnknownModifierType = 0
const UnknownTileType = 0

const NoTileEffects TileEffect = 0

const (
	Passable            TileEffect = 1 << iota // Tile can be passed by units.
	ProvidesFreshWater                         // Tile provides access to fresh water.
	ImpedesMovement                            // Tile slows down unit movement.
	BlocksNavalMovement                        // Tile cannot be navigated by naval units.
	Fertile                                    // Tile is suitable for agriculture.
	RichInWood                                 // Tile is rich in specific resources (minerals, wood, etc.).
	RichInGold                                 // Tile is rich in specific resources (minerals, wood, etc.).
	RichInRock                                 // Tile is rich in specific resources (minerals, wood, etc.).
	RichInFish                                 // Tile is rich in specific resources (minerals, wood, etc.).
	DefensiveBonus                             // Tile provides defensive bonuses to units.
	DamagesUnits                               // Tile causes damage to units over time (e.g., volcanoes, swamps).
	EnhancesMovement                           // Tile enhances movement speed (e.g., roads).
	SupportsNavalUnits                         // Tile supports naval unit movement or docking.
)
