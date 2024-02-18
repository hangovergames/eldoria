// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gamestate

import (
	"github.com/hangovergames/eldoria/internal/server/game/mocks"
	"reflect"
	"testing"
)

func TestNewGameState(t *testing.T) {

	// Setup a mock GameMap for testing
	mockMap := new(mocks.MockGameMap)
	mockMap.On("GetWidth").Return(10)
	mockMap.On("GetHeight").Return(10)
	//tile := new(mocks.MockTile)
	//mockMap.On("GetTile", 10, 10).Return(tile, nil)

	// Call NewGameState with the mock GameMap
	game := NewGameState(mockMap)

	// Verify that the returned *GameState instance has the expected GameMap
	if !reflect.DeepEqual(game.Map, mockMap) {
		t.Errorf("NewGameState() did not initialize GameState with expected GameMap. Got %+v, want %+v", game.Map, mockMap)
	}

	// Additional assertions can be added as necessary, for instance, checking if game.Map is not nil.
	if game.Map.GetWidth() != mockMap.GetWidth() || game.Map.GetHeight() != mockMap.GetHeight() {
		t.Errorf("NewGameState() properties mismatch. Got Width: %d, Height: %d; Want Width: %d, Height: %d",
			game.Map.GetWidth(), game.Map.GetHeight(), mockMap.GetWidth(), mockMap.GetHeight())
	}

	mockMap.AssertExpectations(t)

}
