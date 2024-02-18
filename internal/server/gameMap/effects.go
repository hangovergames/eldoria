// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameMap

import (
	"fmt"
	"github.com/hangovergames/eldoria/internal/server/game"
)

var tileEffectToString = map[game.TileEffect]string{
	game.Passable:            "Passable",
	game.ProvidesFreshWater:  "ProvidesFreshWater",
	game.ImpedesMovement:     "ImpedesMovement",
	game.BlocksNavalMovement: "BlocksNavalMovement",
	game.Fertile:             "Fertile",
	game.RichInWood:          "RichInWood",
	game.RichInGold:          "RichInGold",
	game.RichInRock:          "RichInRock",
	game.RichInFish:          "RichInFish",
	game.DefensiveBonus:      "DefensiveBonus",
	game.DamagesUnits:        "DamagesUnits",
	game.EnhancesMovement:    "EnhancesMovement",
	game.SupportsNavalUnits:  "SupportsNavalUnits",
}

var stringToTileEffect = map[string]game.TileEffect{
	"Passable":            game.Passable,
	"ProvidesFreshWater":  game.ProvidesFreshWater,
	"ImpedesMovement":     game.ImpedesMovement,
	"BlocksNavalMovement": game.BlocksNavalMovement,
	"Fertile":             game.Fertile,
	"RichInWood":          game.RichInWood,
	"RichInGold":          game.RichInGold,
	"RichInRock":          game.RichInRock,
	"RichInFish":          game.RichInFish,
	"DefensiveBonus":      game.DefensiveBonus,
	"DamagesUnits":        game.DamagesUnits,
	"EnhancesMovement":    game.EnhancesMovement,
	"SupportsNavalUnits":  game.SupportsNavalUnits,
}

// TileEffectToString stringifies numeric TileEffect
func TileEffectToString(effect game.TileEffect) string {
	if str, ok := tileEffectToString[effect]; ok {
		return str
	}
	return "Unknown"
}

// StringToTileEffect parses string to numeric TileEffect
func StringToTileEffect(effectStr string) (game.TileEffect, error) {
	if effect, ok := stringToTileEffect[effectStr]; ok {
		return effect, nil
	}
	return 0, fmt.Errorf("unknown effect: %s", effectStr)
}

func CombineTileEffects(effects []game.TileEffect) game.TileEffect {
	var combinedEffect game.TileEffect
	for _, effect := range effects {
		combinedEffect |= effect
	}
	return combinedEffect
}

func StringsToTileEffects(effectStrings []string) ([]game.TileEffect, error) {
	var effects []game.TileEffect
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
