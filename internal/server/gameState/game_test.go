// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameState

import (
	"reflect"
	"testing"

	"github.com/hangovergames/eldoria/internal/server/gameMap"
)

func TestNewGameMap(t *testing.T) {

	// Setup a mock GameMap for testing
	mockMap := gameMap.GameMap{
		// Initialize the GameMap with some mock data. For simplicity, let's assume it's just width and height for now.
		Width:  10,
		Height: 20,
	}

	// Call NewGameState with the mock GameMap
	game := NewGameState(mockMap)

	// Verify that the returned *GameState instance has the expected GameMap
	if !reflect.DeepEqual(game.Map, mockMap) {
		t.Errorf("NewGameState() did not initialize GameState with expected GameMap. Got %+v, want %+v", game.Map, mockMap)
	}

	// Additional assertions can be added as necessary, for instance, checking if game.Map is not nil.
	if game.Map.Width != mockMap.Width || game.Map.Height != mockMap.Height {
		t.Errorf("NewGameState() properties mismatch. Got Width: %d, Height: %d; Want Width: %d, Height: %d",
			game.Map.Width, game.Map.Height, mockMap.Width, mockMap.Height)
	}

}
