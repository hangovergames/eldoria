// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package dtos

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// SpriteSheetDTO represents the configuration for a sprite sheet.
type SpriteSheetDTO struct {
	Name        string `json:"name" yaml:"name"`
	Image       string `json:"image" yaml:"image"`
	TileWidth   int    `json:"tileWidth" yaml:"tileWidth"`
	TileHeight  int    `json:"tileHeight" yaml:"tileHeight"`
	TilesPerRow int    `json:"tilesPerRow" yaml:"tilesPerRow"`
	StartX      int    `json:"startX" yaml:"startX"`
	StartY      int    `json:"startY" yaml:"startY"`
}

// SpriteConfigDTO represents the configuration for mapping sprite names to their sheet and index.
type SpriteConfigDTO struct {
	Name      string `json:"name" yaml:"name"`
	SheetName string `json:"sheetName" yaml:"sheetName"`
	Index     int    `json:"index" yaml:"index"`
}

// SpriteDTO represents the configuration for an individual sprite within a tile.
type SpriteDTO struct {
	Name    string  `json:"name" yaml:"name"`
	XOffset float64 `json:"xOffset" yaml:"xOffset"`
	YOffset float64 `json:"yOffset" yaml:"yOffset"`
}

// TileConfigDTO represents the configuration for a tile, including all its sprites.
type TileConfigDTO struct {
	TileName string      `json:"name" yaml:"name"`
	Sprites  []SpriteDTO `json:"sprites" yaml:"sprites"`
}

// UIConfigDTO represents the top-level configuration structure.
type UIConfigDTO struct {
	SpriteSheets  []SpriteSheetDTO  `json:"spriteSheets" yaml:"spriteSheets"`
	SpriteConfigs []SpriteConfigDTO `json:"sprites" yaml:"sprites"`
	TileConfigs   []TileConfigDTO   `json:"tiles" yaml:"tiles"`
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

	log.Println("Returning: ", config)

	return config
}

// LoadUIConfigDTOFromYAML loads UI configuration from a YAML file.
func LoadUIConfigDTOFromYAML(configPath string) (UIConfigDTO, error) {

	log.Printf("LoadUIConfigDTOFromYAML: %s", configPath)

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

	log.Println("SpriteSheets: ", len(config.SpriteSheets))
	log.Println("SpriteConfigs: ", len(config.SpriteConfigs))
	log.Println("TileConfigs: ", len(config.TileConfigs))

	return config, nil
}
