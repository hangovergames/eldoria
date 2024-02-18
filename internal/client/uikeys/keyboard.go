// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uikeys

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// EbitenKeyboard implements IKeyboard using ebiten's input package.
type EbitenKeyboard struct{}

func (e *EbitenKeyboard) AppendInputChars(inputChars []rune) []rune {
	return ebiten.AppendInputChars(inputChars)
}

func (e *EbitenKeyboard) IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

func (e *EbitenKeyboard) IsKeyJustPressed(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key)
}
