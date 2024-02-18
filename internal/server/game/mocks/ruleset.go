// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import (
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/hangovergames/eldoria/internal/server/game"
	"github.com/stretchr/testify/mock"
)

// MockRuleset is a mock implementation of the IRuleset interface for testing purposes.
type MockRuleset struct {
	mock.Mock
}

// FindTileType simulates finding a tile type by name.
func (m *MockRuleset) FindTileType(tileName string) (game.TileType, bool) {
	args := m.Called(tileName)
	return args.Get(0).(game.TileType), args.Bool(1) // Return whatever is set in your test
}

// FindModifierType simulates finding a modifier type by name.
func (m *MockRuleset) FindModifierType(tileName string) (game.ModifierType, bool) {
	args := m.Called(tileName)
	return args.Get(0).(game.ModifierType), args.Bool(1) // Return whatever is set in your test
}

// CreateTileFromName simulates creating a tile from its name.
func (m *MockRuleset) CreateTileFromName(tileName string) (game.ITile, error) {
	args := m.Called(tileName)
	return args.Get(0).(game.ITile), args.Error(1) // Return whatever is set in your test
}

// IsEffectEnabled simulates checking if an effect is enabled.
func (m *MockRuleset) IsEffectEnabled(effectStr string) bool {
	args := m.Called(effectStr)
	return args.Bool(0) // Return whatever is set in your test
}

// ConvertAndCombineTileEffects simulates converting and combining tile effects.
func (m *MockRuleset) ConvertAndCombineTileEffects(effectStrings []string) (game.TileEffect, error) {
	args := m.Called(effectStrings)
	return args.Get(0).(game.TileEffect), args.Error(1) // Return whatever is set in your test
}

// NewMockRuleset creates an instance of MockRuleset with default values for testing.
func NewMockRuleset() *MockRuleset {
	mockRuleset := &MockRuleset{}
	// Setup default return values for methods if needed
	// Example:
	// mockRuleset.On("FindTileType", "Grassland").Return(TileType(1), true)
	// Adjust according to your test scenarios
	return mockRuleset
}

// GetUI simulates getting the UI configuration.
func (m *MockRuleset) GetUI() dtos.UIConfigDTO {
	args := m.Called()
	// Ensure that you return a UIConfigDTO type.
	// This might require setting up the return value in your test setup.
	return args.Get(0).(dtos.UIConfigDTO) // Cast the return value to dtos.UIConfigDTO
}
