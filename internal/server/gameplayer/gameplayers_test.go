// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameplayer

import (
	"github.com/hangovergames/eldoria/internal/server/game"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		id       game.PlayerID
		wantName string
	}{
		{
			name:     "Test Player 1",
			id:       1,
			wantName: "Test Player 1",
		},
		{
			name:     "Test Player 2",
			id:       2,
			wantName: "Test Player 2",
		},
	}

	// Iterate through test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new Player
			player := NewPlayer(tt.id, tt.name)

			// Verify the ID and Name
			if player.ID != tt.id {
				t.Errorf("NewPlayer() ID = %v, want %v", player.ID, tt.id)
			}
			if player.Name != tt.wantName {
				t.Errorf("NewPlayer() Name = %v, want %v", player.Name, tt.wantName)
			}

			// Verify the Resources map is initialized and empty
			if player.Resources == nil {
				t.Errorf("NewPlayer() Resources map is nil")
			}
			if len(player.Resources) != 0 {
				t.Errorf("NewPlayer() Resources map is not empty")
			}
		})
	}
}
