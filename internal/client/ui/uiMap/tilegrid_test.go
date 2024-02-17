// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uiMap

import (
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/stretchr/testify/assert"
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

func TestLoadTileConfigDTOs(t *testing.T) {

	// Create a mock sprite manager
	mockSpriteManager := new(spriteutils.MockSpriteManager)

	// Create a TileGrid with the mock sprite manager
	width, height := 10, 10
	tileGrid := NewTileGrid(mockSpriteManager, width, height)

	// Define tile configurations to load
	tileConfigs := []dtos.TileConfigDTO{
		{
			TileName: "Grassland",
			Sprites: []dtos.SpriteDTO{
				{Name: "GrasslandSprite", XOffset: 0, YOffset: 0},
			},
		},
		{
			TileName: "Water",
			Sprites: []dtos.SpriteDTO{
				{Name: "WaterSprite", XOffset: 0, YOffset: 0},
				{Name: "WaterDetail", XOffset: 5, YOffset: 5},
			},
		},
	}

	// Load the tile configurations into the TileGrid
	tileGrid.LoadTileConfigDTOs(tileConfigs)

	// Verify that tile configurations are loaded correctly
	// For simplicity, we're checking the existence and count of sprites for one tile type.
	// In a real test, you might want to check more details.
	tileType, exists := tileGrid.nameToID["Grassland"]
	assert.True(t, exists, "Expected Grassland tile to exist")

	config, exists := tileGrid.tileMappings[tileType]
	assert.True(t, exists, "Expected Grassland tile configuration to exist")
	assert.Len(t, config.Sprites, 1, "Expected Grassland to have 1 sprite")

	tileType, exists = tileGrid.nameToID["Water"]
	assert.True(t, exists, "Expected Water tile to exist")

	config, exists = tileGrid.tileMappings[tileType]
	assert.True(t, exists, "Expected Water tile configuration to exist")
	assert.Len(t, config.Sprites, 2, "Expected Water to have 2 sprites")
}
