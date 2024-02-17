// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameMap

import (
	"reflect"
	"testing"
)

func TestTile_Clone(t *testing.T) {

	// Setup an original tile with specific properties
	originalTile := NewTile(1, []ModifierType{2, 3}, Passable|Fertile)

	// Clone the original tile
	clonedTile := originalTile.Clone()

	// Test 1: Check if the cloned tile has the same properties as the original
	if originalTile.Type != clonedTile.Type || !reflect.DeepEqual(originalTile.Modifiers, clonedTile.Modifiers) || originalTile.Effects != clonedTile.Effects {
		t.Errorf("Cloned tile does not match original. Got %+v, want %+v", clonedTile, originalTile)
	}

	// Test 2: Modify the cloned tile's Modifiers and check if the original tile is unaffected
	clonedTile.Modifiers[0] = 99
	if reflect.DeepEqual(originalTile.Modifiers, clonedTile.Modifiers) {
		t.Errorf("Modifying cloned tile's Modifiers affected the original tile. Original %+v, Clone %+v", originalTile, clonedTile)
	}

	// Optional: Test 3: Ensure that the Type and Effects fields are still equal after modifying Modifiers
	if originalTile.Type != clonedTile.Type || originalTile.Effects != clonedTile.Effects {
		t.Errorf("Unexpected change in cloned tile's Type or Effects after modifying Modifiers. Original %+v, Clone %+v", originalTile, clonedTile)
	}

}
