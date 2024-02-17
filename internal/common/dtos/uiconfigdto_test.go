// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package dtos

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestUIConfigDTO(t *testing.T) {

	// Adjust the relative path as necessary
	path, err := filepath.Abs("../../../examples/ui-config-dto.json")
	if err != nil {
		t.Fatalf("Failed to resolve absolute path: %v", err)
	}

	// Read the JSON configuration
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read UI config JSON: %v", err)
	}

	// Unmarshal into UIConfigDTO
	var config UIConfigDTO
	if err := json.Unmarshal(data, &config); err != nil {
		t.Fatalf("Failed to unmarshal UI config JSON: %v", err)
	}

	// Perform assertions
	// Example: Asserting the number of sprite sheets
	expectedSpriteSheets := 2 // Adjust based on your example JSON
	if len(config.SpriteSheets) != expectedSpriteSheets {
		t.Errorf("Expected %d sprite sheets, got %d", expectedSpriteSheets, len(config.SpriteSheets))
	}

	// Further assertions can be added here to validate the contents of TileConfigs, etc.

}
