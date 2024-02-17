// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package spriteutils

import "testing"

// TestNewSpriteIdentifier tests the NewSpriteIdentifier function for correct initialization of SpriteIdentifier instances.
func TestNewSpriteIdentifier(t *testing.T) {
	tests := []struct {
		name      string
		sheetName string
		index     int
	}{
		{
			name:      "BasicTile",
			sheetName: "basic_tiles",
			index:     0,
		},
		{
			name:      "UnitSprite",
			sheetName: "unit_sprites",
			index:     5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			si := NewSpriteIdentifier(tt.sheetName, tt.index)

			if si.SheetName != tt.sheetName {
				t.Errorf("NewSpriteIdentifier().SheetName = %v, want %v", si.SheetName, tt.sheetName)
			}

			if si.Index != tt.index {
				t.Errorf("NewSpriteIdentifier().Index = %v, want %v", si.Index, tt.index)
			}
		})
	}
}
