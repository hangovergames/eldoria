// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package loginscreen

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hangovergames/eldoria/internal/client/drawutils"
	"github.com/hangovergames/eldoria/internal/client/ui"
	"golang.org/x/image/font"
	"image/color"
	"strings"
)

type TextField struct {
	Text                   string
	IsActive               bool
	X, Y                   int // Position of the text field
	Width, Height          int // Dimensions of the text field
	BackspacePressDuration int
	MinLength              int
	MaxLength              int
	CornerRadius           float32 // Radius of the corners
	OnEnter                func()  // Add this line
	FontManager            ui.IFontManager
	FontName               string
	FontFace               font.Face
	TextColor              color.Color
	AllowedChars           string // New field for allowed characters

	PlaceholderFontName string      // Name of the font for placeholder text (optional)
	PlaceholderFontFace font.Face   // Font face for placeholder text (optional)
	Placeholder         string      // Placeholder text
	PlaceholderColor    color.Color // Color of the placeholder text

}

func (tf *TextField) SetFont(fontName string, size float64, dpi float64, clr color.Color) {
	tf.TextColor = clr
	tf.FontName = fontName
	tf.FontFace = tf.FontManager.GetFace(fontName, size, dpi)
}

func (tf *TextField) SetPlaceholderFont(fontName string, size float64, dpi float64, clr color.Color) {
	tf.PlaceholderColor = clr
	tf.PlaceholderFontName = fontName
	tf.PlaceholderFontFace = tf.FontManager.GetFace(fontName, size, dpi)
}

func (tf *TextField) Update() {

	if !tf.IsActive {
		return
	}

	// Create a buffer for input characters with a reasonable initial capacity
	inputChars := make([]rune, 0, 24)

	// Append new characters typed during this frame to the buffer
	inputChars = ebiten.AppendInputChars(inputChars)

	// Append new characters typed
	for _, c := range inputChars {
		if len(tf.Text) < tf.MaxLength && tf.isCharAllowed(c) {
			tf.Text += string(c)
		}
	}

	// Update the backspace press duration
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		tf.BackspacePressDuration++
	} else {
		tf.BackspacePressDuration = 0
	}

	// Handle backspace with repeat delay
	if tf.BackspacePressDuration > 0 && len(tf.Text) > 0 {
		// For simplicity, backspace is handled every 10 frames.
		// Adjust the 10 value for faster or slower repeat.
		if tf.BackspacePressDuration == 1 || (tf.BackspacePressDuration > 10 && tf.BackspacePressDuration%10 == 0) {
			tf.Text = tf.Text[:len(tf.Text)-1]
		}
	}

	// Check if the Enter key is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && len(tf.Text) >= tf.MinLength && tf.OnEnter != nil {
		tf.OnEnter()
	}

}

func (tf *TextField) Draw(screen *ebiten.Image) {

	drawutils.RoundedRect(screen, float32(tf.X), float32(tf.Y), float32(tf.Width), float32(tf.Height), 10, color.RGBA{255, 255, 255, 255})

	if tf.Text == "" && tf.Placeholder != "" {
		drawutils.Text(screen, tf.X+6, tf.Y, tf.Placeholder, tf.PlaceholderFontFace, tf.PlaceholderColor)
	} else {
		drawutils.Text(screen, tf.X+6, tf.Y, tf.Text, tf.FontFace, tf.TextColor)
	}

}

// Helper method to check if a character is allowed
func (tf *TextField) isCharAllowed(c rune) bool {
	if tf.AllowedChars == "" {
		// If AllowedChars is not set, allow all characters
		return true
	}
	// Check if the character is in the AllowedChars string
	return strings.ContainsRune(tf.AllowedChars, c)
}
