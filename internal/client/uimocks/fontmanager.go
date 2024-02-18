// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uimocks

import (
	"github.com/golang/freetype/truetype"
	"github.com/stretchr/testify/mock"
	"golang.org/x/image/font"
)

// MockFontManager is a mock type for the IFontManager interface
type MockFontManager struct {
	mock.Mock
}

// RegisterFont mocks the RegisterFont method
func (m *MockFontManager) RegisterFont(name string, font *truetype.Font) {
	m.Called(name, font)
}

// RegisterFontBytes mocks the RegisterFontBytes method
func (m *MockFontManager) RegisterFontBytes(name string, fontBytes []byte) {
	m.Called(name, fontBytes)
}

// GetFont mocks the GetFont method
func (m *MockFontManager) GetFont(name string) *truetype.Font {
	args := m.Called(name)
	return args.Get(0).(*truetype.Font)
}

// LoadFont mocks the LoadFont method
func (m *MockFontManager) LoadFont(name, filePath string) error {
	args := m.Called(name, filePath)
	return args.Error(0)
}

// GetFace mocks the GetFace method
func (m *MockFontManager) GetFace(name string, size float64, dpi float64) font.Face {
	args := m.Called(name, size, dpi)
	return args.Get(0).(font.Face)
}
