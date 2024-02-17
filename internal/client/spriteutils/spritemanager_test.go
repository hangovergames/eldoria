// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package spriteutils

import (
	"testing"
)

func TestSpriteManager(t *testing.T) {

	spriteManager := NewSpriteManager()

	// Mock sprite sheet registration
	sheet := NewMockSpriteSheet(32, 32, 10, 0, 0)
	spriteManager.RegisterSpriteSheet("testSheet", sheet)

	// Test if the sheet is correctly registered
	if spriteManager.sheets["testSheet"] != sheet {
		t.Errorf("RegisterSpriteSheet failed to register the sheet correctly")
	}

	// Mock sprite name mapping
	spriteManager.MapSpriteName("testSprite", "testSheet", 0)

	// Test if the sprite name is correctly mapped
	identifier, exists := spriteManager.mapping["testSprite"]
	if !exists || identifier.SheetName != "testSheet" || identifier.Index != 0 {
		t.Errorf("MapSpriteName failed to map the sprite name correctly")
	}

	// Assuming a mock or simplified way to verify correctness of the returned *ebiten.Image
	// Since we can't directly verify the *ebiten.Image without graphical operations,
	// we focus on the logical flow: retrieving an existing sprite and handling non-existing ones.
	image := spriteManager.GetSprite("testSprite")
	if image == nil {
		t.Errorf("GetSprite failed to retrieve the sprite correctly")
	}

	// Test retrieval of a non-existing sprite
	if got := spriteManager.GetSprite("nonExistingSprite"); got != nil {
		t.Errorf("GetSprite should return nil for non-existing sprites, got %v", got)
	}

}
