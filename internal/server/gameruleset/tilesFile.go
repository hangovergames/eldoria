// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameruleset

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

// TileConfig represents the configuration for a specific type of tile, including
// its general properties, the effects it has, and the resources it provides.
type TileConfig struct {
	Summary   string                        `yaml:"summary"`   // A brief description of the tile type.
	Effects   []string                      `yaml:"effects"`   // List of effects that this tile imposes.
	Resources map[string]TileResourceConfig `yaml:"resources"` // Definitions of resources available on this tile.
}

// TileResourceConfig defines the characteristics and behaviors of a resource
// as it pertains to its availability and management on a tile.
type TileResourceConfig struct {
	MaxAmount            int  `yaml:"maxAmount"`            // The maximum quantity of the resource that can exist on the tile.
	HarvestRate          Rate `yaml:"harvestRate"`          // The range of amounts of the resource that can be harvested per game turn.
	RestorationRate      Rate `yaml:"restorationRate"`      // The range of amounts the resource can naturally restore per game turn.
	RestorationThreshold int  `yaml:"restorationThreshold"` // The resource amount below which restoration can occur.
	MaxPopulation        int  `yaml:"maxPopulation"`        // The maximum number of population units that can be assigned to harvest the resource.
}

// Rate defines a range with a minimum and maximum value, used to specify
// variable rates for resource harvesting and restoration.
type Rate struct {
	Min int `yaml:"min"` // The minimum value in the range.
	Max int `yaml:"max"` // The maximum value in the range.
}

// ModifierConfig represents the effects and resource modifiers associated with each tile or modifier.
type ModifierConfig struct {
	Summary string   `yaml:"summary"` // A brief description of the modifier.
	Effects []string `yaml:"effects"` // List of effects that this modifier imposes.
}

type TilesFile struct {
	Tiles     map[string]TileConfig     `yaml:"tiles"`
	Modifiers map[string]ModifierConfig `yaml:"modifiers"`
}

func LoadTilesFile(path string) (TilesFile, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return TilesFile{}, fmt.Errorf("error reading YAML file: %v", err)
	}

	var ruleset TilesFile
	err = yaml.Unmarshal(data, &ruleset)
	if err != nil {
		return TilesFile{}, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return ruleset, nil
}
