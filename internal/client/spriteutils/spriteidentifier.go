// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package spriteutils

// SpriteIdentifier uniquely identifies a sprite within a collection of sprite sheets.
type SpriteIdentifier struct {
	SheetName string // The name of the sprite sheet.
	Index     int    // The index of the sprite within the sheet.
}

// NewSpriteIdentifier creates a new instance of SpriteIdentifier.
func NewSpriteIdentifier(
	sheetName string,
	index int,
) *SpriteIdentifier {
	return &SpriteIdentifier{
		SheetName: sheetName,
		Index:     index,
	}
}
