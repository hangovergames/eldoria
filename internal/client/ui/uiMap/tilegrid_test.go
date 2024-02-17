// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uiMap

import (
	"testing"

	"github.com/hangovergames/eldoria/internal/client/spriteutils"
)

// TestDefineTileConfig tests the DefineTileConfig method of TileGrid.
func TestDefineTileConfig(t *testing.T) {

	// Initialize TileGrid with a mock sprite manager and dimensions.
	mockSpriteManager := &spriteutils.MockSpriteManager{}
	uiMap := NewTileGrid(mockSpriteManager, 10, 10) // Assume a constructor NewTileGrid exists.

	// Define a tile configuration.
	tileType := uint(0)
	tileName := "Ocean"
	name := "ocean"
	xOffset := float64(15)
	yOffset := float64(15)

	// First call to DefineTileConfig.
	uiMap.DefineTileConfig(tileName, name, xOffset, yOffset)

	// Verify that the tile configuration was created.
	if config, exists := uiMap.tileMappings[tileType]; !exists || len(config.Sprites) != 1 {
		t.Errorf("Expected 1 sprite configuration, got %d", len(config.Sprites))
	}

	// Second call to DefineTileConfig with the same tileType but different offsets.
	uiMap.DefineTileConfig(tileName, "oceanEdge", 0, 0)

	// Verify that the second sprite configuration was appended.
	if config, exists := uiMap.tileMappings[tileType]; !exists || len(config.Sprites) != 2 {
		t.Errorf("Expected 2 sprite configurations after second call, got %d", len(config.Sprites))
	}

	// Validate the details of the appended configuration.
	if config, _ := uiMap.tileMappings[tileType]; config.Sprites[1].Name != "oceanEdge" {
		t.Errorf("Expected second sprite name to be 'oceanEdge', got '%s'", config.Sprites[1].Name)
	}

}
