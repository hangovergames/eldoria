// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/eldoria/internal/client/ui"
	"github.com/hangovergames/eldoria/internal/client/uimocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/hangovergames/eldoria/internal/client/uimap"
)

func TestGameUIDraw(t *testing.T) {

	// Create a mock TileGrid
	mockGrid := new(uimap.MockTileGrid)

	mockScreen := new(uimocks.MockScreen)
	// Set expectation
	mockScreen.On("Draw", mock.AnythingOfType("*ebiten.Image")).Once()

	// Create a GameUI instance with the mock grid
	gameUI := NewGameUI(800, 600, mockGrid)

	gameUI.RegisterScreen("Mock", func() ui.IScreen {
		return mockScreen
	})

	gameUI.SetCurrentScreen("Mock")

	// Create a dummy ebiten.Image (could be just a 1x1 image for simplicity)
	screen := ebiten.NewImage(1, 1)

	// Call the Draw method
	gameUI.Draw(screen)

	// Assert that Draw was called on the mockGrid
	mockScreen.AssertExpectations(t)
	mockGrid.AssertExpectations(t)

}
