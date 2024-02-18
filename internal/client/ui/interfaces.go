// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package ui

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"golang.org/x/image/font"
)

// ITileGrid defines the interface for a tile grid.
type ITileGrid interface {
	Draw(screen *ebiten.Image, tileSizeX, tileSizeY int)
	SetTile(x, y int, tileNames ...string)
	GetTile(x, y int) ([]string, bool)
	LoadTileConfigDTOs(tileConfigs []dtos.TileConfigDTO)
}

type IScreen interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)
}

type IGame interface {
	GetMap() ITileGrid
	GetScreenWidth() int
	GetScreenHeight() int
}

type IImageManager interface {
	GetImage(name string) *ebiten.Image
}

type IFontManager interface {

	// RegisterFont associates a *truetype.Font with a name.
	RegisterFont(name string, font *truetype.Font)

	// RegisterFontBytes registers a font from a byte slice.
	RegisterFontBytes(name string, fontBytes []byte)

	// GetFont retrieves a *truetype.Font by name.
	GetFont(name string) *truetype.Font

	// LoadFont loads a font from a file and registers it under the given name.
	LoadFont(name, filePath string) error

	// GetFace returns a font.Face for a registered font with the specified size and DPI.
	GetFace(name string, size float64, dpi float64) font.Face
}

// IKeyboard defines an interface for keyboard input handling.
type IKeyboard interface {
	AppendInputChars(inputChars []rune) []rune
	IsKeyPressed(key ebiten.Key) bool
	IsKeyJustPressed(key ebiten.Key) bool
}

// ISpriteIdentifier defines the behavior for sprite identification and retrieval.
type ISpriteIdentifier interface {

	// GetSheetName returns the name of the sprite sheet.
	GetSheetName() string

	// GetIndex returns the index of the sprite within the sheet.
	GetIndex() int
}

// ISpriteSheet defines the methods for interacting with sprite sheets.
type ISpriteSheet interface {
	// SubImage extracts a sub-image from the sprite sheet based on the tile index.
	SubImage(index int) *ebiten.Image
}

// ISpriteManager defines the behavior for managing sprites.
type ISpriteManager interface {
	GetSprite(name string) *ebiten.Image
	RegisterSpriteSheet(name string, sheet ISpriteSheet)
	MapSpriteName(name string, sheetName string, index int)
	LoadSpriteSheetDTOs(spriteSheets []dtos.SpriteSheetDTO)
	LoadSpriteConfigDTOs(spriteConfigs []dtos.SpriteConfigDTO)
}
