// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package ruleset

import (
	"github.com/hangovergames/eldoria/internal/testutils"
	"testing"
)

// TestLoadTilesFile_Success tests loading a well-formed YAML file.
func TestLoadTilesFile_Success(t *testing.T) {
	content := `
tiles:
  Tile1:
    summary: "Summary for tile 1"
    effects:
      - effect1
modifiers:
  Modifier1:
    summary: "Summary for modifier 1"
    effects:
      - effect1
`
	filename, cleanup := testutils.CreateTempYAMLFile(t, content)
	defer cleanup()

	ruleset, err := LoadTilesFile(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(ruleset.Tiles) != 1 || len(ruleset.Modifiers) != 1 {
		t.Errorf("Unexpected number of tiles or modifiers loaded")
	}
}

// TestLoadTilesFile_FileNotFound tests the behavior when the YAML file does not exist.
func TestLoadTilesFile_FileNotFound(t *testing.T) {
	_, err := LoadTilesFile("nonexistent.yaml")
	if err == nil {
		t.Errorf("Expected an error for a nonexistent file, got nil")
	}
}
