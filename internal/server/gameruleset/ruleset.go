// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameruleset

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/hangovergames/eldoria/internal/server/game"
	"github.com/hangovergames/eldoria/internal/server/gamemap"
)

type Ruleset struct {
	EnabledTiles     []string // TilesEnabled index of this array is also the numeric tile ID on the map
	EnabledModifiers []string
	EnabledEffects   []string
	Tiles            TilesFile
	Resources        ResourcesFile
	Effects          EffectsFile
	UI               dtos.UIConfigDTO
}

// LoadRuleset combines loading of RulesetFile and TilesFile into a single Ruleset struct.
func LoadRuleset(rulesetPath string) (Ruleset, error) {

	// Check if the path is already absolute
	if !filepath.IsAbs(rulesetPath) {
		// If not, prepend the working directory
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get working directory: %v", err)
		}
		rulesetPath = filepath.Join(wd, rulesetPath)
	}

	// Load the basic gameRuleset
	rulesetFilePath := filepath.Join(rulesetPath, "ruleset.yml")
	rulesetFile, err := LoadRulesetFile(rulesetFilePath)
	if err != nil {
		return Ruleset{}, fmt.Errorf("failed to load %s: %w", rulesetFilePath, err)
	}

	// Load tiles
	tilesFilePath := filepath.Join(rulesetPath, "tiles.yml")
	tilesFile, err := LoadTilesFile(tilesFilePath)
	if err != nil {
		return Ruleset{}, fmt.Errorf("failed to load %s: %w", tilesFilePath, err)
	}

	// Load resources
	resourcesFilePath := filepath.Join(rulesetPath, "resources.yml")
	resourcesFile, err := LoadResourcesFile(resourcesFilePath)
	if err != nil {
		return Ruleset{}, fmt.Errorf("failed to load %s: %w", resourcesFilePath, err)
	}

	// Load effects
	effectsFilePath := filepath.Join(rulesetPath, "effects.yml")
	effectsFile, err := LoadEffectsFile(effectsFilePath)
	if err != nil {
		return Ruleset{}, fmt.Errorf("failed to load %s: %w", effectsFilePath, err)
	}

	// Load UI config
	uiFilePath := filepath.Join(rulesetPath, "ui.yml")
	uiFile, err := dtos.LoadUIConfigDTOFromYAML(uiFilePath)
	if err != nil {
		return Ruleset{}, fmt.Errorf("failed to load %s: %w", uiFilePath, err)
	}
	log.Println("ui = ", uiFile)

	// Combine the loaded data into a single Ruleset struct
	ruleset := Ruleset{
		EnabledTiles:     rulesetFile.Tiles,
		EnabledModifiers: rulesetFile.Modifiers,
		EnabledEffects:   rulesetFile.Effects,
		Tiles:            tilesFile,
		Resources:        resourcesFile,
		Effects:          effectsFile,
		UI:               uiFile,
	}

	return ruleset, nil
}

// FindTileType finds the index of a given tile name in the Ruleset.EnabledTiles slice,
// which represents the TileType.
// It returns the index and a boolean indicating whether the tile was found.
func (r *Ruleset) FindTileType(tileName string) (game.TileType, bool) {
	for index, name := range r.EnabledTiles {
		if name == tileName {
			return game.TileType(index), true
		}
	}
	return game.UnknownTileType, false
}

// FindModifierType finds the index of a given tile modifier name in the
// Ruleset.EnabledModifiers slice, which represents the ModifierType.
// It returns the index and a boolean indicating whether the tile modifier was found.
func (r *Ruleset) FindModifierType(tileName string) (game.ModifierType, bool) {
	for index, name := range r.EnabledModifiers {
		if name == tileName {
			return game.ModifierType(index), true
		}
	}
	return game.UnknownModifierType, false // Return -1 when not found
}

// CreateTileFromName creates a Tile based on its name using the definitions in
// the Ruleset. It ensures that only enabled effects are applied. The tile will
// not have any modifiers set.
func (r *Ruleset) CreateTileFromName(
	tileName string,
) (game.ITile, error) {

	tileType, found := r.FindTileType(tileName)
	if !found {
		return nil, fmt.Errorf("tile '%s' not found in enabled tiles", tileName)
	}

	tileConfig, exists := r.Tiles.Tiles[tileName]
	if !exists {
		return nil, fmt.Errorf("tile '%s' does not have a configuration", tileName)
	}

	combinedEffects, err := r.ConvertAndCombineTileEffects(tileConfig.Effects)
	if err != nil {
		fmt.Println("Error combining effects:", err)
		return nil, fmt.Errorf("tile '%s' failed to setup effects", tileName)
	}

	var tile = gamemap.NewTile(tileType, []game.ModifierType{}, combinedEffects)

	return &tile, nil
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
func (r *Ruleset) ConvertAndCombineTileEffects(effectStrings []string) (game.TileEffect, error) {

	var effects []game.TileEffect
	for _, effectStr := range effectStrings {
		if r.IsEffectEnabled(effectStr) {
			effect, err := gamemap.StringToTileEffect(effectStr)
			if err != nil {
				// Decide how you want to handle the error.
				// For this example, let's just return the error along with what we've combined so far.
				// You could also choose to log the error and continue without returning.
				return gamemap.CombineTileEffects(effects), fmt.Errorf("error converting string '%s' to TileEffect: %w", effectStr, err)
			}
			effects = append(effects, effect)
		}
	}

	// Combine all enabled effects for this tile.
	combinedEffects := gamemap.CombineTileEffects(effects)
	return combinedEffects, nil
}

func (r *Ruleset) GetUI() dtos.UIConfigDTO {
	return r.UI
}
