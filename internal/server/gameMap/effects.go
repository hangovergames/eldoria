// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameMap

import "fmt"

type TileEffect uint

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

var tileEffectToString = map[TileEffect]string{
	Passable:            "Passable",
	ProvidesFreshWater:  "ProvidesFreshWater",
	ImpedesMovement:     "ImpedesMovement",
	BlocksNavalMovement: "BlocksNavalMovement",
	Fertile:             "Fertile",
	RichInWood:          "RichInWood",
	RichInGold:          "RichInGold",
	RichInRock:          "RichInRock",
	RichInFish:          "RichInFish",
	DefensiveBonus:      "DefensiveBonus",
	DamagesUnits:        "DamagesUnits",
	EnhancesMovement:    "EnhancesMovement",
	SupportsNavalUnits:  "SupportsNavalUnits",
}

var stringToTileEffect = map[string]TileEffect{
	"Passable":            Passable,
	"ProvidesFreshWater":  ProvidesFreshWater,
	"ImpedesMovement":     ImpedesMovement,
	"BlocksNavalMovement": BlocksNavalMovement,
	"Fertile":             Fertile,
	"RichInWood":          RichInWood,
	"RichInGold":          RichInGold,
	"RichInRock":          RichInRock,
	"RichInFish":          RichInFish,
	"DefensiveBonus":      DefensiveBonus,
	"DamagesUnits":        DamagesUnits,
	"EnhancesMovement":    EnhancesMovement,
	"SupportsNavalUnits":  SupportsNavalUnits,
}

// TileEffectToString stringifies numeric TileEffect
func TileEffectToString(effect TileEffect) string {
	if str, ok := tileEffectToString[effect]; ok {
		return str
	}
	return "Unknown"
}

// StringToTileEffect parses string to numeric TileEffect
func StringToTileEffect(effectStr string) (TileEffect, error) {
	if effect, ok := stringToTileEffect[effectStr]; ok {
		return effect, nil
	}
	return 0, fmt.Errorf("unknown effect: %s", effectStr)
}

func CombineTileEffects(effects []TileEffect) TileEffect {
	var combinedEffect TileEffect
	for _, effect := range effects {
		combinedEffect |= effect
	}
	return combinedEffect
}

func StringsToTileEffects(effectStrings []string) ([]TileEffect, error) {
	var effects []TileEffect
	for _, effectStr := range effectStrings {
		effect, err := StringToTileEffect(effectStr)
		if err != nil {
			// If an error occurs, you can decide to either return the error immediately,
			// or continue processing other strings. Here, we return immediately.
			return nil, err
		}
		effects = append(effects, effect)
	}
	return effects, nil
}

// HasTileEffect Function to check if a tile effect includes a specific effect
func (t TileEffect) HasTileEffect(effect TileEffect) bool {
	return t&effect != 0
}
