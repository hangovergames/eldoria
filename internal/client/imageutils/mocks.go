// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package imageutils

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// MockImageManager simulates the IImageManager interface for testing purposes.
type MockImageManager struct {
	Images map[string]*ebiten.Image
}

// NewMockImageManager creates a new instance of MockImageManager.
func NewMockImageManager() *MockImageManager {
	return &MockImageManager{
		Images: make(map[string]*ebiten.Image),
	}
}

// GetImage retrieves a mock *ebiten.Image by name.
func (m *MockImageManager) GetImage(name string) *ebiten.Image {
	if img, exists := m.Images[name]; exists {
		return img
	}
	// Optionally return a default mock image or nil if the image is not found
	return nil
}
