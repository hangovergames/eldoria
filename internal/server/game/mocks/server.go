// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import (
	"github.com/hangovergames/eldoria/internal/server/game"
	"github.com/stretchr/testify/mock"
)

// MockServer is a mock implementation of the IServer interface for testing purposes.
type MockServer struct {
	mock.Mock
	Address string
	Ruleset game.IRuleset
	State   game.IGameState
}

// Start simulates starting the server. It doesn't do anything in the mock.
func (m *MockServer) Start() error {
	args := m.Called()
	return args.Error(0) // Return nil or an error depending on what's set in your test
}

// SetupRoutes simulates setting up routes. It doesn't do anything in the mock.
func (m *MockServer) SetupRoutes() {
	m.Called()
}

// GetAddress returns a mock server address.
func (m *MockServer) GetAddress() string {
	args := m.Called()
	return args.String(0) // Return the Address field or a test-specific address
}

// GetState returns a mock server address.
func (m *MockServer) GetState() game.IGameState {
	args := m.Called()
	return args.Get(0).(game.IGameState) // Ensure your tests set this up correctly
}

// GetRuleset returns a mock ruleset.
func (m *MockServer) GetRuleset() game.IRuleset {
	args := m.Called()
	return args.Get(0).(game.IRuleset) // Ensure your tests set this up correctly
}

// NewMockServer creates an instance of MockServer with default values for testing.
func NewMockServer() *MockServer {
	mockServer := &MockServer{}
	// Setup default return values for methods if needed
	mockServer.On("Start").Return(nil)
	mockServer.On("GetAddress").Return("http://localhost:8080")
	// Ensure you setup the mock ruleset here or in your tests directly
	return mockServer
}
