// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameruleset

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type RulesetFile struct {
	Tiles     []string `yaml:"tiles"`
	Modifiers []string `yaml:"modifiers"`
	Effects   []string `yaml:"effects"`
}

func LoadRulesetFile(path string) (RulesetFile, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return RulesetFile{}, fmt.Errorf("error reading YAML file: %v", err)
	}

	var ruleset RulesetFile
	err = yaml.Unmarshal(data, &ruleset)
	if err != nil {
		return RulesetFile{}, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return ruleset, nil
}
