// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gamemap

import (
	"github.com/hangovergames/eldoria/internal/server/game"
	"reflect"
	"testing"
)

var DeepOcean game.ITile = NewTile(0, []game.ModifierType{}, game.NoTileEffects)
var Grassland game.ITile = NewTile(1, []game.ModifierType{}, game.NoTileEffects)

func TestNewGameMap(t *testing.T) {
	gm := NewGameMap(10, 5, DeepOcean)
	if gm.Width != 10 || gm.Height != 5 {
		t.Errorf("NewGameMap dimensions incorrect, got: width %d, height %d", gm.Width, gm.Height)
	}
	for _, row := range gm.Tiles {
		for _, tile := range row {
			if tile.GetType() != DeepOcean.GetType() {
				t.Errorf("NewGameMap default tile incorrect, expected: 0, got: %d", tile)
			}
		}
	}
}

func TestGetSetTile(t *testing.T) {
	gm := NewGameMap(10, 5, DeepOcean)

	// Set a tile and then get it
	err := gm.SetTile(3, 2, Grassland)
	if err != nil {
		t.Fatalf("SetTile returned an error: %v", err)
	}
	tile, err := gm.GetTile(3, 2)
	if err != nil {
		t.Fatalf("GetTile returned an error: %v", err)
	}
	if tile.GetType() != Grassland.GetType() {
		t.Errorf("Expected tile type Grassland, got: %d", tile.GetType())
	}

	// Test out of bounds
	_, err = gm.GetTile(-1, 0)
	if err == nil {
		t.Error("Expected error for out of bounds coordinates, got nil")
	}
}

func TestGetTilesInArea(t *testing.T) {

	gm := NewGameMap(10, 5, DeepOcean)
	gm.SetTile(1, 1, Grassland) // Set a specific tile to make the test meaningful

	area, err := gm.GetTilesInArea(0, 0, 2, 2)
	if err != nil {
		t.Fatalf("GetTilesInArea returned an error: %v", err)
	}
	expected := [][]game.ITile{
		{DeepOcean, DeepOcean, DeepOcean},
		{DeepOcean, Grassland, DeepOcean},
		{DeepOcean, DeepOcean, DeepOcean},
	}
	if !reflect.DeepEqual(area, expected) {
		t.Errorf("Expected area to be %+v, got %+v", expected, area)
	}

	// Test invalid coordinates
	_, err = gm.GetTilesInArea(-1, 0, 10, 10)
	if err == nil {
		t.Error("Expected error for invalid coordinates, got nil")
	}
}
