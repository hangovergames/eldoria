// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameruleset

import (
	"github.com/hangovergames/eldoria/internal/common/testutils"
	"testing"
)

func TestLoadRulesetFile_Success(t *testing.T) {
	// Create a temporary YAML file with valid content
	content := `
tiles:
  - tile1
  - tile2
modifiers:
  - modifier1
effects:
  - effect1
`
	filename, cleanup := testutils.CreateTempYAMLFile(t, content)
	defer cleanup()

	// Attempt to load the gameRuleset
	ruleset, err := LoadRulesetFile(filename)
	if err != nil {
		t.Errorf("Ruleset failed to load")
	}

	// Verify the content of the loaded gameRuleset
	if len(ruleset.Tiles) != 2 || len(ruleset.Modifiers) != 1 || len(ruleset.Effects) != 1 {
		t.Errorf("Ruleset content did not match expected values")
	}
}

func TestLoadRulesetFile_FileNotFound(t *testing.T) {
	_, err := LoadRulesetFile("nonexistent.yaml")
	if err == nil {
		t.Errorf("Expected an error for nonexistent file, got nil")
	}
}
