// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/eldoria/internal/client/ui/uiMap"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGameUIDraw(t *testing.T) {

	// Create a mock TileGrid
	mockGrid := new(uiMap.MockTileGrid)
	// Set expectation
	mockGrid.On("Draw", mock.Anything, 30, 30).Once()

	// Create a GameUI instance with the mock grid
	gameUI := NewGameUI(800, 600, mockGrid)

	// Create a dummy ebiten.Image (could be just a 1x1 image for simplicity)
	screen := ebiten.NewImage(1, 1)

	// Call the Draw method
	gameUI.Draw(screen)

	// Assert that Draw was called on the mockGrid
	mockGrid.AssertExpectations(t)

}
