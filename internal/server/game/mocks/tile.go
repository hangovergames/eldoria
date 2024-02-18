// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import (
	"github.com/hangovergames/eldoria/internal/server/game"
	"github.com/stretchr/testify/mock"
)

// MockTile is a mock type for the ITile interface
type MockTile struct {
	mock.Mock
}

// GetType mocks the GetType method
func (m *MockTile) GetType() game.TileType {
	args := m.Called()
	return args.Get(0).(game.TileType)
}

// GetModifiers mocks the GetModifiers method
func (m *MockTile) GetModifiers() []game.ModifierType {
	args := m.Called()
	return args.Get(0).([]game.ModifierType)
}

// GetEffects mocks the GetEffects method
func (m *MockTile) GetEffects() game.TileEffect {
	args := m.Called()
	return args.Get(0).(game.TileEffect)
}

// Clone mocks the Clone method
func (m *MockTile) Clone() game.ITile {
	args := m.Called()
	return args.Get(0).(game.ITile)
}
