// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package spriteutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/eldoria/internal/client/imageutils"
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"log"
)

// ISpriteManager defines the behavior for managing sprites.
type ISpriteManager interface {
	GetSprite(name string) *ebiten.Image
	RegisterSpriteSheet(name string, sheet *SpriteSheet)
	MapSpriteName(name string, sheetName string, index int)
	LoadSpriteSheetDTOs(spriteSheets []dtos.SpriteSheetDTO)
	LoadSpriteConfigDTOs(spriteConfigs []dtos.SpriteConfigDTO)
}

// SpriteManager manages multiple SpriteSheets and provides an easy way to retrieve sprites by name.
type SpriteManager struct {
	sheets       map[string]*SpriteSheet      // Map of sprite sheet names to SpriteSheet instances.
	mapping      map[string]*SpriteIdentifier // Map of sprite names to their identifiers (sheet name and index).
	imageManager imageutils.IImageManager     // Image manager
}

// NewSpriteManager creates a new instance of SpriteManager.
func NewSpriteManager(imageManager imageutils.IImageManager) *SpriteManager {
	return &SpriteManager{
		sheets:       make(map[string]*SpriteSheet),
		mapping:      make(map[string]*SpriteIdentifier),
		imageManager: imageManager,
	}
}

// RegisterSpriteSheet associates a SpriteSheet with a name within the manager.
func (sm *SpriteManager) RegisterSpriteSheet(name string, sheet *SpriteSheet) {
	sm.sheets[name] = sheet
}

// MapSpriteName maps a logical sprite name to its identifier in a sprite sheet.
func (sm *SpriteManager) MapSpriteName(
	name string, // name the sprite name.
	sheetName string, // sheetName The name of the sprite sheet.
	index int, // index The index of the sprite within the sheet.
) {
	sm.mapping[name] = NewSpriteIdentifier(sheetName, index)
}

// GetSprite retrieves a sprite by its logical name.
func (sm *SpriteManager) GetSprite(name string) *ebiten.Image {
	identifier, exists := sm.mapping[name]
	if !exists {
		return nil
	}
	sheet, exists := sm.sheets[identifier.SheetName]
	if !exists {
		return nil
	}
	return sheet.SubImage(identifier.Index)
}

// LoadSpriteSheetDTOs loads sprite sheets defined in UIConfigDTO.
func (sm *SpriteManager) LoadSpriteSheetDTOs(spriteSheets []dtos.SpriteSheetDTO) {
	for _, sheetDTO := range spriteSheets {
		img := sm.imageManager.GetImage(sheetDTO.Image)
		if img == nil {
			log.Printf("Image not found for sprite sheet: %s", sheetDTO.Image)
			continue
		}
		sheet := NewSpriteSheet(img, sheetDTO.TileWidth, sheetDTO.TileHeight, sheetDTO.TilesPerRow, sheetDTO.StartX, sheetDTO.StartY)
		sm.RegisterSpriteSheet(sheetDTO.Name, sheet)
	}
}

// LoadSpriteConfigDTOs loads sprite configurations defined in the UIConfigDTO into the SpriteManager.
func (sm *SpriteManager) LoadSpriteConfigDTOs(spriteConfigs []dtos.SpriteConfigDTO) {
	for _, config := range spriteConfigs {
		sm.MapSpriteName(config.Name, config.SheetName, config.Index)
	}
}
