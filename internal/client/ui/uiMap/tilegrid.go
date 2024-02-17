// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uiMap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/eldoria/internal/client/spriteutils"
)

type SpriteConfig struct {
	Name    string  // Name of the sprite.
	XOffset float64 // Horizontal offset for the sprite.
	YOffset float64 // Vertical offset for the sprite.
}

type TileConfig struct {
	Sprites []SpriteConfig // Configuration for sprites that compose this tile.
}

// TileGrid represents the UI component for drawing a map.
type TileGrid struct {
	spriteManager spriteutils.ISpriteManager
	width, height int                 // Map size in tiles.
	tileMappings  map[uint]TileConfig // Mapping from tile type to visual configuration.
	mapGrid       [][]uint            // Grid defining the tile type at each position.
	nameToID      map[string]uint     // Mapping from tile names to uint identifiers.
	nextID        uint                // Next available uint identifier for a new tile type.
}

// NewTileGrid creates a new instance of TileGrid. Each tile will be 0.
func NewTileGrid(spriteManager spriteutils.ISpriteManager, width, height int) *TileGrid {

	// Initialize the map grid with the specified dimensions.
	mapGrid := make([][]uint, height)
	for i := range mapGrid {
		mapGrid[i] = make([]uint, width) // Initialize each row with the specified width.
	}

	// Return a new TileGrid instance.
	return &TileGrid{
		spriteManager: spriteManager,
		width:         width,
		height:        height,
		tileMappings:  make(map[uint]TileConfig), // Initialize the tile mappings.
		mapGrid:       mapGrid,                   // Assign the initialized map grid.
		nameToID:      make(map[string]uint),
		nextID:        0,
	}

}

// DefineTileConfig associates a tile type with its visual configuration.
func (um *TileGrid) DefineTileConfig(
	tileName string, // tileName Numeric presentation of a tile
	spriteName string, // spriteName Name of the sprite.
	xOffset float64, // xOffset Horizontal offset for the sprite.
	yOffset float64, // yOffset Vertical offset for the sprite.
) {

	// Check if the tileName already has an assigned ID
	id, exists := um.nameToID[tileName]
	if !exists {
		// Assign a new ID and update the mapping
		id = um.nextID
		um.nextID++
		um.nameToID[tileName] = id
	}

	// Proceed with the configuration using the resolved ID
	if config, exists := um.tileMappings[id]; exists {
		config.Sprites = append(config.Sprites, SpriteConfig{Name: spriteName, XOffset: xOffset, YOffset: yOffset})
		um.tileMappings[id] = config
	} else {
		um.tileMappings[id] = TileConfig{
			Sprites: []SpriteConfig{{Name: spriteName, XOffset: xOffset, YOffset: yOffset}},
		}
	}

}

// Draw renders the map to the given Ebiten image reference.
func (um *TileGrid) Draw(
	screen *ebiten.Image,
	tileSizeX, tileSizeY int,
) {
	for y, row := range um.mapGrid {
		for x, tileType := range row {

			config, exists := um.tileMappings[tileType]
			if !exists {
				continue // Skip undefined tiles
			}

			for _, spriteConfig := range config.Sprites {

				sprite := um.spriteManager.GetSprite(spriteConfig.Name)

				// Skip missing sprites
				if sprite == nil {
					continue
				}

				opts := &ebiten.DrawImageOptions{}

				// Apply scale if necessary
				//opts.GeoM.Scale(float64(tileSizeX)/float64(sprite.Bounds().Dx()), float64(tileSizeY)/float64(sprite.Bounds().Dy()))

				// Apply the specified offsets for each sprite
				opts.GeoM.Translate(float64(x*tileSizeX)+spriteConfig.XOffset, float64(y*tileSizeY)+spriteConfig.YOffset)

				screen.DrawImage(sprite, opts)

			}
		}
	}
}

// SetTile sets the tile type at the given position.
func (tg *TileGrid) SetTile(x, y int, tileName string) {
	id, exists := tg.nameToID[tileName]
	if !exists {
		return // Optionally handle error or log warning
	}
	if x >= 0 && x < tg.width && y >= 0 && y < tg.height {
		tg.mapGrid[y][x] = id
	}
}

// GetTile returns the tile type at the given position.
// The boolean return value indicates whether the position is within the grid bounds.
func (tg *TileGrid) GetTile(x, y int) (string, bool) {
	if x < 0 || x >= tg.width || y < 0 || y >= tg.height {
		return "", false
	}
	id := tg.mapGrid[y][x]
	for name, idMapping := range tg.nameToID {
		if id == idMapping {
			return name, true
		}
	}
	return "", false // Tile not found or is the default tile
}
