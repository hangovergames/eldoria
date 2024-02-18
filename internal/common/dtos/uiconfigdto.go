// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package dtos

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
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

// SpriteConfigDTO represents the configuration for mapping sprite names to their sheet and index.
type SpriteConfigDTO struct {
	Name      string `json:"name"`
	SheetName string `json:"sheetName"`
	Index     int    `json:"index"`
}

// SpriteDTO represents the configuration for an individual sprite within a tile.
type SpriteDTO struct {
	Name    string  `json:"name"`
	XOffset float64 `json:"xOffset"`
	YOffset float64 `json:"yOffset"`
}

// TileConfigDTO represents the configuration for a tile, including all its sprites.
type TileConfigDTO struct {
	TileName string      `json:"name"`
	Sprites  []SpriteDTO `json:"sprites"`
}

// UIConfigDTO represents the top-level configuration structure.
type UIConfigDTO struct {
	SpriteSheets  []SpriteSheetDTO  `json:"spriteSheets"`
	SpriteConfigs []SpriteConfigDTO `json:"sprites"`
	TileConfigs   []TileConfigDTO   `json:"tiles"`
}

func LoadUIConfigDTO(relativeConfigPath string) UIConfigDTO {

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	// Construct the path to the configuration file
	configPath := filepath.Join(wd, relativeConfigPath)

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

	log.Println("Returning: ", config)

	return config
}

// LoadUIConfigDTOFromYAML loads UI configuration from a YAML file.
func LoadUIConfigDTOFromYAML(configPath string) (UIConfigDTO, error) {

	// Read the YAML file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return UIConfigDTO{}, fmt.Errorf("failed to read configuration file: %v", err)
	}

	// Unmarshal the YAML into the UIConfigDTO struct
	var config UIConfigDTO
	if err := yaml.Unmarshal(data, &config); err != nil {
		return UIConfigDTO{}, fmt.Errorf("failed to parse configuration YAML: %v", err)
	}

	log.Println("Returning: ", config)

	return config, nil
}
