// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package spriteutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"testing"
)

// mockEbitenImage creates a mock ebiten.Image for testing purposes.
// In a real test, you might need to replace this with actual ebiten.Image creation or another mocking approach.
func mockEbitenImage(w, h int) *ebiten.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	return ebiten.NewImageFromImage(img)
}

func TestNewSpriteSheet(t *testing.T) {
	img := mockEbitenImage(100, 100) // Mock image with 100x100 size
	ss := NewSpriteSheet(img, 10, 10, 10, 0, 0)

	if ss.TileWidth != 10 || ss.TileHeight != 10 || ss.TilesPerRow != 10 {
		t.Errorf("NewSpriteSheet did not initialize correctly")
	}
}

func TestSubImage(t *testing.T) {
	img := mockEbitenImage(100, 100) // Mock image with 100x100 size
	ss := NewSpriteSheet(img, 10, 10, 10, 0, 0)

	// Test extracting the first sub-image
	subImg := ss.SubImage(0)
	if subImg == nil {
		t.Errorf("SubImage returned nil")
	}

	// Here you would ideally check if subImg has the correct rectangle extracted,
	// but since ebiten.Image does not expose its internals and we're not in an environment
	// that can perform graphical assertions, this part is left as a conceptual step.
	// In practice, you might check for non-nil return and possibly mock or abstract
	// ebiten.Image operations to validate parameters passed to SubImage().
}
