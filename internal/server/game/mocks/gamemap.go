// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import (
	"github.com/hangovergames/eldoria/internal/server/game"
	"github.com/stretchr/testify/mock"
)

// MockGameMap is a mock type for the IGameMap interface
type MockGameMap struct {
	mock.Mock
}

// GetTile mocks the GetTile method
func (m *MockGameMap) GetTile(x, y int) (game.ITile, error) {
	args := m.Called(x, y)
	return args.Get(0).(game.ITile), args.Error(1)
}

// SetTile mocks the SetTile method
func (m *MockGameMap) SetTile(x, y int, newTile game.ITile) error {
	args := m.Called(x, y, newTile)
	return args.Error(0)
}

// GetTilesInArea mocks the GetTilesInArea method
func (m *MockGameMap) GetTilesInArea(x, y, x2, y2 int) ([][]game.ITile, error) {
	args := m.Called(x, y, x2, y2)
	return args.Get(0).([][]game.ITile), args.Error(1)
}

func (m *MockGameMap) GetWidth() int {
	args := m.Called()
	return args.Get(0).(int)
}

func (m *MockGameMap) GetHeight() int {
	args := m.Called()
	return args.Get(0).(int)
}
