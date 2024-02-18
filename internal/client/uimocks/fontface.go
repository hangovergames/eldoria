// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uimocks

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
)

// MockFontFace is a mock type that implements the font.Face interface.
type MockFontFace struct {
	mock.Mock
}

func (m *MockFontFace) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockFontFace) Glyph(dot fixed.Point26_6, r rune) (image.Rectangle, image.Image, image.Point, fixed.Int26_6, bool) {
	args := m.Called(dot, r)
	return args.Get(0).(image.Rectangle), args.Get(1).(image.Image), args.Get(2).(image.Point), args.Get(3).(fixed.Int26_6), args.Bool(4)
}

func (m *MockFontFace) GlyphBounds(r rune) (fixed.Rectangle26_6, fixed.Int26_6, bool) {
	args := m.Called(r)
	return args.Get(0).(fixed.Rectangle26_6), args.Get(1).(fixed.Int26_6), args.Bool(2)
}

func (m *MockFontFace) GlyphAdvance(r rune) (fixed.Int26_6, bool) {
	args := m.Called(r)
	return args.Get(0).(fixed.Int26_6), args.Bool(1)
}

func (m *MockFontFace) Kern(r0, r1 rune) fixed.Int26_6 {
	args := m.Called(r0, r1)
	return args.Get(0).(fixed.Int26_6)
}

func (m *MockFontFace) Metrics() font.Metrics {
	args := m.Called()
	return args.Get(0).(font.Metrics)
}
