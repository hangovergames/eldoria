// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/mock"
)

// MockScreen is a mock type for the IScreen interface
type MockScreen struct {
	mock.Mock
}

// Update mocks the Update method
func (m *MockScreen) Update() error {
	args := m.Called()
	return args.Error(0)
}

// Draw mocks the Draw method
func (m *MockScreen) Draw(screen *ebiten.Image) {
	m.Called(screen)
}

// Layout mocks the Layout method
func (m *MockScreen) Layout(outsideWidth, outsideHeight int) (int, int) {
	args := m.Called(outsideWidth, outsideHeight)
	return args.Int(0), args.Int(1)
}
