// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import (
	"github.com/hangovergames/eldoria/internal/server/game"
	"github.com/stretchr/testify/mock"
)

// MockGameState is a mock type for the IGameState interface.
type MockGameState struct {
	mock.Mock
}

// GetMap mocks the GetMap method.
func (m *MockGameState) GetMap() game.IGameMap {
	args := m.Called()
	return args.Get(0).(game.IGameMap) // Ensure to return the correct type.
}

// GetPlayers mocks the GetPlayers method.
func (m *MockGameState) GetPlayers() []game.IPlayer {
	args := m.Called()
	return args.Get(0).([]game.IPlayer) // Ensure to return the correct type.
}

// FindPlayer mocks the GetPlayer method.
func (m *MockGameState) FindPlayer(name string) (game.IPlayer, error) {
	args := m.Called(name)
	return args.Get(0).(game.IPlayer), args.Get(1).(error)
}
