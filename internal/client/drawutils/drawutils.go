// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package drawutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hangovergames/eldoria/internal/common/assets"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"log"
	"math"
)

func RoundedRect(screen *ebiten.Image, x, y, width, height, r float32, clr color.Color) {

	path := vector.Path{}

	// Top left corner
	path.MoveTo(x+r, y)
	path.Arc(x+r, y+r, r, -math.Pi, -math.Pi/2, vector.Clockwise)

	// Top right corner
	path.LineTo(x+width-r, y)
	path.Arc(x+width-r, y+r, r, -math.Pi/2, 0, vector.Clockwise)

	// Bottom right corner
	path.LineTo(x+width, y+height-r)
	path.Arc(x+width-r, y+height-r, r, 0, math.Pi/2, vector.Clockwise)

	// Bottom left corner
	path.LineTo(x+r, y+height)
	path.Arc(x+r, y+height-r, r, math.Pi/2, math.Pi, vector.Clockwise)

	// Close the path
	path.LineTo(x, y+r)

	// Create vertices and indices for filling the path
	vertices, indices := path.AppendVerticesAndIndicesForFilling(nil, nil)

	// Get the RGBA components from the color
	r32, g32, b32, a32 := clr.RGBA()
	r, g, b, a := float32(r32)/0x101, float32(g32)/0x101, float32(b32)/0x101, float32(a32)/0x101

	// Set color and source position for each vertex
	for i := range vertices {
		vertices[i].SrcX = 0
		vertices[i].SrcY = 0
		vertices[i].ColorR = r / 0xff
		vertices[i].ColorG = g / 0xff
		vertices[i].ColorB = b / 0xff
		vertices[i].ColorA = a / 0xff
	}

	// Draw the rounded rectangle on the screen
	screen.DrawTriangles(vertices, indices, assets.WhiteSubImage, nil)

}

// DrawText directly draws text onto an ebiten.Image at the specified location.
// This function requires the caller to provide a font.Face for text styling.
func Text(
	screen *ebiten.Image,
	x, y int,
	text string,
	fontFace font.Face,
	clr color.Color,
) {
	if fontFace == nil {
		log.Println("Font face not defined for DrawText")
		return
	}

	// Create a font drawer and configure it for text rendering
	drawer := font.Drawer{
		Dst:  screen,
		Src:  image.NewUniform(clr),
		Face: fontFace,
		Dot: fixed.Point26_6{
			X: fixed.I(x),                                         // Starting X position
			Y: fixed.I(y + int(fontFace.Metrics().Ascent.Ceil())), // Adjust Y position based on font ascent
		},
	}

	// Draw the text string onto the screen
	drawer.DrawString(text)
}
