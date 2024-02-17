// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package dtos

import (
	"encoding/json"
	"log"
	"os"
)

// SpriteSheetDTO represents the configuration for a sprite sheet.
type SpriteSheetDTO struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	TileWidth   int    `json:"tileWidth"`
	TileHeight  int    `json:"tileHeight"`
	TilesPerRow int    `json:"tilesPerRow"`
	StartX      int    `json:"startX"`
	StartY      int    `json:"startY"`
}

// SpriteDTO represents the configuration for an individual sprite within a tile.
type SpriteDTO struct {
	Name    string  `json:"name"`
	XOffset float64 `json:"xOffset"`
	YOffset float64 `json:"yOffset"`
}

// TileConfigDTO represents the configuration for a tile, including all its sprites.
type TileConfigDTO struct {
	TileName string      `json:"tileName"`
	Sprites  []SpriteDTO `json:"sprites"`
}

// UIConfigDTO represents the top-level configuration structure.
type UIConfigDTO struct {
	SpriteSheets []SpriteSheetDTO `json:"spriteSheets"`
	TileConfigs  []TileConfigDTO  `json:"tileConfigs"`
}

func LoadUIConfigDTO(configPath string) UIConfigDTO {

	// Read the JSON file
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	// Unmarshal the JSON into the ConfigurationDTO struct
	var config UIConfigDTO
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse configuration JSON: %v", err)
	}

	return config
}
