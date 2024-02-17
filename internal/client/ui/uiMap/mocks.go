// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package uiMap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/stretchr/testify/mock"
)

// MockTileGrid mocks the TileGrid for testing purposes.
type MockTileGrid struct {
	mock.Mock
}

// Draw simulates drawing the tile grid.
func (m *MockTileGrid) Draw(screen *ebiten.Image, tileSizeX, tileSizeY int) {
	m.Called(screen, tileSizeX, tileSizeY)
}

// SetTile simulates setting a tile at a given position.
func (m *MockTileGrid) SetTile(x, y int, tileName string) {
	m.Called(x, y, tileName)
}

// GetTile simulates getting a tile at a given position.
func (m *MockTileGrid) GetTile(x, y int) (string, bool) {
	args := m.Called(x, y)
	return args.String(0), args.Bool(1)
}

// Implement the LoadTileConfigDTOs method to satisfy the ITileGrid interface.
func (m *MockTileGrid) LoadTileConfigDTOs(tileConfigs []dtos.TileConfigDTO) {
	m.Called(tileConfigs)
}
