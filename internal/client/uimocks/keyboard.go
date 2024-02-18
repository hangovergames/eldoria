// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uimocks

import "github.com/hajimehoshi/ebiten/v2"

// MockKeyboard is a mock implementation of IKeyboard for testing.
type MockKeyboard struct {
	InputChars      []rune
	PressedKeys     map[ebiten.Key]bool
	JustPressedKeys map[ebiten.Key]bool
}

func NewMockKeyboard() *MockKeyboard {
	return &MockKeyboard{
		JustPressedKeys: make(map[ebiten.Key]bool),
	}
}

// AppendInputChars returns a predefined set of input characters.
func (m *MockKeyboard) AppendInputChars(inputChars []rune) []rune {
	return append(inputChars, m.InputChars...)
}

// IsKeyPressed returns whether the specified key was "just pressed" based on predefined state.
func (m *MockKeyboard) IsKeyPressed(key ebiten.Key) bool {
	return m.PressedKeys[key]
}

// IsKeyJustPressed returns whether the specified key was "just pressed" based on predefined state.
func (m *MockKeyboard) IsKeyJustPressed(key ebiten.Key) bool {
	return m.JustPressedKeys[key]
}
