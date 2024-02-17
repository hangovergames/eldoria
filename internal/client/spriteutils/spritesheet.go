// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package spriteutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// SpriteSheet represents a sprite sheet with properties that define how its images are structured.
type SpriteSheet struct {
	Image       *ebiten.Image // The entire sprite sheet image.
	TileWidth   int           // Width of a single tile in the sheet.
	TileHeight  int           // Height of a single tile in the sheet.
	TilesPerRow int           // Number of tiles per row in the sprite sheet.
	StartX      int           // The starting X coordinate of the tile grid within the sprite sheet.
	StartY      int           // The starting Y coordinate of the tile grid within the sprite sheet.
}

// NewSpriteSheet creates a new SpriteSheet instance with the specified properties.
func NewSpriteSheet(img *ebiten.Image, tileWidth, tileHeight, tilesPerRow, startX, startY int) *SpriteSheet {
	return &SpriteSheet{
		Image:       img,
		TileWidth:   tileWidth,
		TileHeight:  tileHeight,
		TilesPerRow: tilesPerRow,
		StartX:      startX,
		StartY:      startY,
	}
}

// SubImage extracts a sub-image from the sprite sheet based on the tile index.
func (ss *SpriteSheet) SubImage(index int) *ebiten.Image {
	x := ss.StartX + (index%ss.TilesPerRow)*ss.TileWidth
	y := ss.StartY + (index/ss.TilesPerRow)*ss.TileHeight
	return ss.Image.SubImage(image.Rect(x, y, x+ss.TileWidth, y+ss.TileHeight)).(*ebiten.Image)
}
