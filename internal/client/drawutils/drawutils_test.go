// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package drawutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"testing"
)

// Example font face for testing. In a real test, you would use a mock or a simple font.Face implementation.
type fakeFontFace struct{}

func (f fakeFontFace) Close() error { return nil }
func (f fakeFontFace) Glyph(dot fixed.Point26_6, rune rune) (image.Rectangle, image.Image, image.Point, fixed.Int26_6, bool) {
	return image.Rectangle{}, nil, image.Point{}, 0, false
}
func (f fakeFontFace) GlyphBounds(rune rune) (fixed.Rectangle26_6, fixed.Int26_6, bool) {
	return fixed.Rectangle26_6{}, 0, false
}
func (f fakeFontFace) GlyphAdvance(rune rune) (fixed.Int26_6, bool) { return 0, false }
func (f fakeFontFace) Kern(r0, r1 rune) fixed.Int26_6               { return 0 }
func (f fakeFontFace) Metrics() font.Metrics                        { return font.Metrics{} }

func TestRoundedRect(t *testing.T) {
	screen := ebiten.NewImage(100, 100) // Create a blank image
	clr := color.RGBA{255, 0, 0, 255}   // Red color

	RoundedRect(screen, 10, 10, 80, 80, 10, clr)

	// Inspect specific pixels if necessary. For simplicity, this is omitted.
	// You would need to check that the function modified the image as expected.
}

func TestText(t *testing.T) {
	screen := ebiten.NewImage(100, 100) // Create a blank image
	clr := color.RGBA{0, 255, 0, 255}   // Green color
	text := "Test"
	fontFace := fakeFontFace{}

	Text(screen, 10, 10, text, fontFace, clr)

	// Similar to RoundedRect, check specific pixels to ensure the text was drawn.
	// This is simplified and omitted here.
}
