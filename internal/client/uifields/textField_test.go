// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uifields

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/eldoria/internal/client/uimocks"
	"testing"
)

func TestTextField_AllowedChars(t *testing.T) {

	// Create a mock keyboard and set up the input simulation.
	mockKeyboard := uimocks.NewMockKeyboard()
	mockKeyboard.InputChars = []rune{'a'}                 // Simulate 'a' input for the first call.
	mockKeyboard.JustPressedKeys[ebiten.KeyEnter] = false // Simulate that Enter key is not pressed.

	fontManager := new(uimocks.MockFontManager) // Assuming you have a mock font manager
	fontFace := new(uimocks.MockFontFace)

	textField := TextField{
		X:            0,
		Y:            0,
		Width:        200,
		Height:       50,
		MaxLength:    10,
		IsActive:     true,
		FontFace:     fontFace,
		FontManager:  fontManager,
		AllowedChars: "abc",
		Keyboard:     mockKeyboard,
	}

	// Call the Update method, which should now use the mock keyboard for input.
	textField.Update()

	// Assert that 'a' was allowed and added to the text.
	if textField.Text != "a" {
		t.Errorf("Expected 'a' to be allowed, but got '%s'", textField.Text)
	}

	// Reset mockKeyboard.InputChars to simulate 'd' input and call Update again.
	mockKeyboard.InputChars = []rune{'d'}
	textField.Update()

	// Assert that 'd' was disallowed and the text remains unchanged.
	if textField.Text != "a" {
		t.Errorf("Expected 'd' to be disallowed, but Text was modified to '%s'", textField.Text)
	}

}
