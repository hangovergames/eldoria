// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package spriteutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// NewMockSpriteSheet creates a mock SpriteSheet for testing.
func NewMockSpriteSheet(tileWidth, tileHeight, tilesPerRow, startX, startY int) *SpriteSheet {

	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	ebitenImg := ebiten.NewImageFromImage(img)

	return &SpriteSheet{
		Image:       ebitenImg,
		TileWidth:   tileWidth,
		TileHeight:  tileHeight,
		TilesPerRow: tilesPerRow,
		StartX:      startX,
		StartY:      startY,
	}
}
