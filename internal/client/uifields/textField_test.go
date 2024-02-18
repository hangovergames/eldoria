// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uifields

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/eldoria/internal/client/uimocks"
	"github.com/stretchr/testify/assert"
	"image/color"
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

// TestPlaceholderVisibility checks if the placeholder is displayed when the text field is empty and not displayed when the text field is not empty.
func TestPlaceholderVisibility(t *testing.T) {
	fontManager := new(uimocks.MockFontManager) // Assuming you have a mock font manager
	fontFace := new(uimocks.MockFontFace)

	fontManager.On("GetFace", "some-font", float64(12), float64(72)).Return(fontFace)

	mockKeyboard := uimocks.NewMockKeyboard()
	textField := TextField{
		IsActive:         true,
		FontFace:         fontFace,
		FontManager:      fontManager,
		Placeholder:      "Enter text...",
		PlaceholderColor: color.Gray{Y: 128},
		Text:             "",
		Keyboard:         mockKeyboard,
	}
	textField.SetPlaceholderFont("some-font", 12, 72, color.Gray{Y: 128})

	// When text is empty, placeholder should be visible.
	textField.Update()
	assert.Equal(t, "", textField.Text, "Text should be empty")
	// Placeholder visibility is typically checked during rendering, you might simulate drawing and check for placeholder.

	// Simulate text input
	mockKeyboard.InputChars = []rune{'h', 'e', 'l', 'l', 'o'}
	textField.Update()
	assert.Equal(t, "hello", textField.Text, "Text should be updated with input")

	// Now, placeholder should not be visible.
	// Again, this would be checked in the rendering logic which isn't directly testable without rendering.

}

// TestMaxLength checks if the text field does not exceed the specified maximum length.
func TestMaxLength(t *testing.T) {
	fontManager := new(uimocks.MockFontManager) // Assuming you have a mock font manager
	fontFace := new(uimocks.MockFontFace)
	mockKeyboard := uimocks.NewMockKeyboard()
	textField := TextField{
		IsActive:    true,
		FontFace:    fontFace,
		FontManager: fontManager,
		MaxLength:   5,
		Keyboard:    mockKeyboard,
	}

	mockKeyboard.InputChars = []rune{'h', 'e', 'l', 'l', 'o', 'w', 'o', 'r', 'l', 'd'}
	textField.Update()

	assert.Equal(t, "hello", textField.Text, "Text should not exceed MaxLength")
}

// TestAllowedChars verifies that only allowed characters are added to the text field.
func TestAllowedChars(t *testing.T) {
	fontManager := new(uimocks.MockFontManager) // Assuming you have a mock font manager
	fontFace := new(uimocks.MockFontFace)
	mockKeyboard := uimocks.NewMockKeyboard()
	textField := TextField{
		IsActive:     true,
		FontFace:     fontFace,
		FontManager:  fontManager,
		AllowedChars: "abc",
		Keyboard:     mockKeyboard,
	}

	mockKeyboard.InputChars = []rune{'a', 'b', 'c', 'd', 'e'}
	textField.Update()

	assert.Equal(t, "abc", textField.Text, "Text should only contain allowed characters")
}

// TestOnEnterCallback checks if the OnEnter callback is executed when the Enter key is pressed.
func TestOnEnterCallback(t *testing.T) {
	fontManager := new(uimocks.MockFontManager) // Assuming you have a mock font manager
	fontFace := new(uimocks.MockFontFace)
	mockKeyboard := uimocks.NewMockKeyboard()
	called := false
	textField := TextField{
		IsActive:    true,
		OnEnter:     func() { called = true },
		Keyboard:    mockKeyboard,
		FontFace:    fontFace,
		FontManager: fontManager,
	}

	// Simulate Enter key press
	mockKeyboard.JustPressedKeys[ebiten.KeyEnter] = true
	textField.Update()

	assert.True(t, called, "OnEnter callback should be called")
}
