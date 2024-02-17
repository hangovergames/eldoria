// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameRuleset

import (
	"gopkg.in/yaml.v2"
	"os"
)

// AllResourcesKeyword This is a special key which affects all resources if
// defined in Resources map
const AllResourcesKeyword = "all"

// ResourceConfig represents a basic configuration for a resource,
// containing only a summary description.
type ResourceConfig struct {
	Summary string `yaml:"summary"`
}

// ResourcesFile maps resource names to their configurations.
// It's designed to load the simplified resources definition.
type ResourcesFile struct {
	Resources map[string]ResourceConfig `yaml:"resources"`
}

// LoadResourcesFile loads resource definitions from a YAML file.
func LoadResourcesFile(filePath string) (ResourcesFile, error) {
	var resourcesFile ResourcesFile

	data, err := os.ReadFile(filePath)
	if err != nil {
		return resourcesFile, err
	}

	err = yaml.Unmarshal(data, &resourcesFile)
	if err != nil {
		return resourcesFile, err
	}

	return resourcesFile, nil
}
