// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameRuleset

import (
	"fmt"
	"github.com/hangovergames/eldoria/internal/server/gameMap"
	"path/filepath"
)

type Ruleset struct {
	EnabledTiles     []string // TilesEnabled index of this array is also the numeric tile ID on the map
	EnabledModifiers []string
	EnabledEffects   []string
	Tiles            TilesFile
	Resources        ResourcesFile
	Effects          EffectsFile
}

// LoadRuleset combines loading of RulesetFile and TilesFile into a single Ruleset struct.
func LoadRuleset(rulesetPath string) (Ruleset, error) {

	// Load the basic gameRuleset
	rulesetFilePath := filepath.Join(rulesetPath, "gameRuleset.yml")
	rulesetFile, err := LoadRulesetFile(rulesetFilePath)
	if err != nil {
		return Ruleset{}, fmt.Errorf("failed to load gameRuleset.yml: %w", err)
	}

	// Load tiles
	tilesFilePath := filepath.Join(rulesetPath, "tiles.yml")
	tilesFile, err := LoadTilesFile(tilesFilePath)
	if err != nil {
		return Ruleset{}, fmt.Errorf("failed to load tiles.yml: %w", err)
	}

	// Load resources
	resourcesFilePath := filepath.Join(rulesetPath, "resources.yml")
	resourcesFile, err := LoadResourcesFile(resourcesFilePath)
	if err != nil {
		return Ruleset{}, fmt.Errorf("failed to load resources.yml: %w", err)
	}

	// Load effects
	effectsFilePath := filepath.Join(rulesetPath, "effects.yml")
	effectsFile, err := LoadEffectsFile(effectsFilePath)
	if err != nil {
		return Ruleset{}, fmt.Errorf("failed to load effects.yml: %w", err)
	}

	// Combine the loaded data into a single Ruleset struct
	ruleset := Ruleset{
		EnabledTiles:     rulesetFile.Tiles,
		EnabledModifiers: rulesetFile.Modifiers,
		EnabledEffects:   rulesetFile.Effects,
		Tiles:            tilesFile,
		Resources:        resourcesFile,
		Effects:          effectsFile,
	}

	return ruleset, nil
}

// FindTileType finds the index of a given tile name in the Ruleset.EnabledTiles slice,
// which represents the TileType.
// It returns the index and a boolean indicating whether the tile was found.
func (r *Ruleset) FindTileType(tileName string) (gameMap.TileType, bool) {
	for index, name := range r.EnabledTiles {
		if name == tileName {
			return gameMap.TileType(index), true
		}
	}
	return gameMap.UnknownTileType, false // Return -1 when not found
}

// FindModifierType finds the index of a given tile modifier name in the
// Ruleset.EnabledModifiers slice, which represents the ModifierType.
// It returns the index and a boolean indicating whether the tile modifier was found.
func (r *Ruleset) FindModifierType(tileName string) (gameMap.ModifierType, bool) {
	for index, name := range r.EnabledModifiers {
		if name == tileName {
			return gameMap.ModifierType(index), true
		}
	}
	return gameMap.UnknownModifierType, false // Return -1 when not found
}

// CreateTileFromName creates a Tile based on its name using the definitions in
// the Ruleset. It ensures that only enabled effects are applied. The tile will
// not have any modifiers set.
func (r *Ruleset) CreateTileFromName(
	tileName string,
) (gameMap.Tile, error) {

	tileType, found := r.FindTileType(tileName)
	if !found {
		return gameMap.Tile{}, fmt.Errorf("tile '%s' not found in enabled tiles", tileName)
	}

	tileConfig, exists := r.Tiles.Tiles[tileName]
	if !exists {
		return gameMap.Tile{}, fmt.Errorf("tile '%s' does not have a configuration", tileName)
	}

	combinedEffects, err := r.ConvertAndCombineTileEffects(tileConfig.Effects)
	if err != nil {
		fmt.Println("Error combining effects:", err)
		return gameMap.Tile{}, fmt.Errorf("tile '%s' failed to setup effects", tileName)
	}

	return gameMap.NewTile(tileType, []gameMap.ModifierType{}, combinedEffects), nil
}

// IsEffectEnabled checks if a given effect string is in the list of enabled effects.
func (r *Ruleset) IsEffectEnabled(effectStr string) bool {
	for _, enabledEffect := range r.EnabledEffects {
		if enabledEffect == effectStr {
			return true
		}
	}
	return false
}

// ConvertAndCombineTileEffects converts a slice of effect strings to TileEffect enums,
// filters out any effects not enabled in the Ruleset, combines them, and returns the combined TileEffect.
// If an error occurs during conversion, it can either ignore the error and continue, or return the combined effects up to that point along with the error.
func (r *Ruleset) ConvertAndCombineTileEffects(effectStrs []string) (gameMap.TileEffect, error) {

	var effects []gameMap.TileEffect
	for _, effectStr := range effectStrs {
		if r.IsEffectEnabled(effectStr) {
			effect, err := gameMap.StringToTileEffect(effectStr)
			if err != nil {
				// Decide how you want to handle the error.
				// For this example, let's just return the error along with what we've combined so far.
				// You could also choose to log the error and continue without returning.
				return gameMap.CombineTileEffects(effects), fmt.Errorf("error converting string '%s' to TileEffect: %w", effectStr, err)
			}
			effects = append(effects, effect)
		}
	}

	// Combine all enabled effects for this tile.
	combinedEffects := gameMap.CombineTileEffects(effects)
	return combinedEffects, nil
}
