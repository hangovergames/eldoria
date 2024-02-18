// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameMap

import (
	"github.com/hangovergames/eldoria/internal/server/game"
	"reflect"
	"testing"
)

func TestTile_Clone(t *testing.T) {

	// Setup an original tile with specific properties
	originalTile := NewTile(1, []game.ModifierType{2, 3}, game.Passable|game.Fertile)

	// Clone the original tile
	clonedTile := originalTile.Clone()

	// Test 1: Check if the cloned tile has the same properties as the original
	if originalTile.Type != clonedTile.GetType() || !reflect.DeepEqual(originalTile.Modifiers, clonedTile.GetModifiers()) || originalTile.Effects != clonedTile.GetEffects() {
		t.Errorf("Cloned tile does not match original. Got %+v, want %+v", clonedTile, originalTile)
	}

	// Test 2: Modify the cloned tile's Modifiers and check if the original tile is unaffected
	clonedTile.GetModifiers()[0] = 99
	if reflect.DeepEqual(originalTile.Modifiers, clonedTile.GetModifiers()) {
		t.Errorf("Modifying cloned tile's Modifiers affected the original tile. Original %+v, Clone %+v", originalTile, clonedTile)
	}

	// Optional: Test 3: Ensure that the Type and Effects fields are still equal after modifying Modifiers
	if originalTile.Type != clonedTile.GetType() || originalTile.Effects != clonedTile.GetEffects() {
		t.Errorf("Unexpected change in cloned tile's Type or Effects after modifying Modifiers. Original %+v, Clone %+v", originalTile, clonedTile)
	}

}
