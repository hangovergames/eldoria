// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameRuleset

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

// EffectConfig represents the configuration for an effect, including its summary
// and its impact on resources.
type EffectConfig struct {
	Summary   string                        `yaml:"summary"`   // A brief description of the effect.
	Resources map[string]TileResourceConfig `yaml:"resources"` // Specifies how the effect impacts resources.
}

// EffectsFile represents the container for all effects defined in the YAML.
type EffectsFile struct {
	Effects map[string]EffectConfig `yaml:"effects"`
}

func LoadEffectsFile(path string) (EffectsFile, error) {
	var effectsFile EffectsFile

	data, err := os.ReadFile(path)
	if err != nil {
		return EffectsFile{}, fmt.Errorf("error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(data, &effectsFile)
	if err != nil {
		return EffectsFile{}, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return effectsFile, nil
}
