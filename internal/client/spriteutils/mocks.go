// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package spriteutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"image/color"
)

// MockSpriteManager simulates SpriteManager for testing purposes.
type MockSpriteManager struct {
	// Optional: Add a map to store fake or test sprites by name.
	fakeSprites map[string]*ebiten.Image
	// Track calls for verification in tests.
	CalledGetSprite            map[string]bool
	CalledLoadSpriteSheetDTOs  bool
	CalledLoadSpriteConfigDTOs bool
}

// NewMockSpriteManager creates a new instance of MockSpriteManager.
func NewMockSpriteManager() *MockSpriteManager {
	return &MockSpriteManager{
		fakeSprites:     make(map[string]*ebiten.Image),
		CalledGetSprite: make(map[string]bool),
	}
}

// GetSprite simulates retrieving a sprite by name. Can return a fake sprite for testing.
func (m *MockSpriteManager) GetSprite(name string) *ebiten.Image {
	m.CalledGetSprite[name] = true // Track that GetSprite was called with this name.

	// Check if a fake sprite is stored for this name.
	if sprite, exists := m.fakeSprites[name]; exists {
		return sprite
	}

	// For simplicity in this example, return a plain image. In real tests, adjust as needed.
	img := ebiten.NewImage(10, 10)       // Small, placeholder image.
	img.Fill(color.RGBA{R: 255, A: 255}) // Optional: Fill with a distinctive color.
	return img
}

// AddFakeSprite allows tests to add a fake sprite to the mock.
func (m *MockSpriteManager) AddFakeSprite(name string, sprite *ebiten.Image) {
	m.fakeSprites[name] = sprite
}

// Implement stubs for other methods if needed for tracking or simulation.
func (sm *MockSpriteManager) RegisterSpriteSheet(name string, sheet *SpriteSheet) {
	// Optional: Implement if you need to simulate or track this call.
}

func (sm *MockSpriteManager) MapSpriteName(name string, sheetName string, index int) {
	// Optional: Implement if you need to simulate or track this call.
}

// Implement the LoadSpriteSheetDTOs method for the MockSpriteManager.
func (m *MockSpriteManager) LoadSpriteSheetDTOs(spriteSheets []dtos.SpriteSheetDTO) {
	// Mark that the method was called. No need to implement logic for mocking.
	m.CalledLoadSpriteSheetDTOs = true
}

// Implement the LoadSpriteConfigDTOs method for the MockSpriteManager.
func (m *MockSpriteManager) LoadSpriteConfigDTOs(spriteConfigs []dtos.SpriteConfigDTO) {
	// Mark that the method was called. No need to implement logic for mocking.
	m.CalledLoadSpriteConfigDTOs = true
}
