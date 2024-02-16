// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameMap

import "fmt"

type GameMap struct {
	Tiles         [][]Tile
	Width, Height int
}

func NewGameMap(width, height int, defaultTileType TileType, modifiers ...ModifierType) *GameMap {
	tiles := make([][]Tile, height)
	for i := range tiles {
		tiles[i] = make([]Tile, width)
		for j := range tiles[i] {
			tiles[i][j] = NewTile(defaultTileType, modifiers...)
		}
	}
	return &GameMap{
		Tiles:  tiles,
		Width:  width,
		Height: height,
	}
}

// GetTile returns the ID of the tile at the given coordinates.
func (gm *GameMap) GetTile(x, y int) (Tile, error) {
	if x < 0 || y < 0 || x >= gm.Width || y >= gm.Height {
		return NewTile(UnknownTileType), fmt.Errorf("coordinates out of bounds")
	}
	return gm.Tiles[y][x], nil
}

// SetTile sets the ID of the tile at the given coordinates.
func (gm *GameMap) SetTile(x, y int, tileID TileType, modifiers ...ModifierType) error {
	if x < 0 || y < 0 || x >= gm.Width || y >= gm.Height {
		return fmt.Errorf("coordinates out of bounds")
	}
	gm.Tiles[y][x] = NewTile(tileID, modifiers...)
	return nil
}

// GetTilesInArea returns a 2D slice of tile IDs within the specified rectangular area.
func (gm *GameMap) GetTilesInArea(x, y, x2, y2 int) ([][]Tile, error) {

	// Validate coordinates
	if x < 0 || y < 0 || x2 >= gm.Width || y2 >= gm.Height || x > x2 || y > y2 {
		return nil, fmt.Errorf("invalid coordinates")
	}

	// Calculate the width and height of the area
	width := x2 - x + 1
	height := y2 - y + 1

	// Initialize the slice to hold the tiles
	areaTiles := make([][]Tile, height)
	for i := range areaTiles {
		areaTiles[i] = make([]Tile, width)
	}

	// Populate the slice with tile IDs
	for i := y; i <= y2; i++ {
		for j := x; j <= x2; j++ {
			areaTiles[i-y][j-x] = gm.Tiles[i][j]
		}
	}

	return areaTiles, nil
}
