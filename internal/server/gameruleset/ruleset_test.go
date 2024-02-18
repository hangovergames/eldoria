// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameruleset

import (
	"github.com/hangovergames/eldoria/internal/common/testutils"
	"github.com/hangovergames/eldoria/internal/server/game"
	"log"
	"testing"
)

func TestLoadRuleset(t *testing.T) {

	// Create a temporary directory for the test environment
	dirPath, cleanupDir := testutils.CreateTempDir(t)
	defer cleanupDir()

	log.Printf("dirPath = %s", dirPath)

	// Create ruleset.yml
	rulesetContent := `tiles:
- Unknown
- ShallowOcean
- Grassland
modifiers:
- River
- Road
- CoralReefs
effects:
- BlocksNavalMovement
- SupportsNavalUnits
- Passable
`
	_, cleanupFile := testutils.CreateFileInDir(t, dirPath, "ruleset.yml", rulesetContent)
	defer cleanupFile()

	// Create tiles.yml
	tilesContent := `tiles:
  Grassland:
    summary: Grassland
    effects:
    - Fertile
    - Passable
    
  Forest:
    summary: Forest
    effects:
    - ImpedesMovement
    - RichInWood

modifiers:
  River:
    summary: River
    effects:
    - EnhancesMovement
    - ProvidesFreshWater
  Road:
    summary: Road
    effects:
    - EnhancesMovement
  CoralReefs:
    summary: CoralReefs
    effects:
    - BlocksNavalMovement
`
	_, cleanupTilesFile := testutils.CreateFileInDir(t, dirPath, "tiles.yml", tilesContent)
	defer cleanupTilesFile()

	// Create resources.yml
	resourcesContent := `
resources:
  Gold:
    summary: Gold summary
  Wood:
    summary: Wood summary
`
	_, cleanupResourcesFile := testutils.CreateFileInDir(t, dirPath, "resources.yml", resourcesContent)
	defer cleanupResourcesFile()

	// Create effects.yml
	effectsContent := `
effects:
  HarvestGold:
    summary: Gold summary
  HarvestWood:
    summary: Wood summary
`
	_, cleanupEffectsFile := testutils.CreateFileInDir(t, dirPath, "effects.yml", effectsContent)
	defer cleanupEffectsFile()

	// Create ui.yml
	uiContent := `
spriteSheets:

- name: Tiles
  image: "freeciv/data/trident/tiles.png"
  tileWidth: 30
  tileHeight: 30
  tilesPerRow: 20
  startX: 0
  startY: 0

- name: OceanTiles
  image: "freeciv/data/trident/tiles.png"
  tileWidth: 15
  tileHeight: 15
  tilesPerRow: 32
  startX: 0
  startY: 210

sprites:

- name: ShallowOcean
  sheetName: OceanTiles
  index: 0

- name: DeepOcean
  sheetName: OceanTiles
  index: 10

- name: Grassland
  sheetName: Tiles
  index: 2

tiles:

- name: DeepOcean
  sprites:
  - name: DeepOcean
    xOffset: 0
    yOffset: 0
  - name: DeepOcean
    xOffset: 15
    yOffset: 0
  - name: DeepOcean
    xOffset: 0
    yOffset: 15
  - name: DeepOcean
    xOffset: 15
    yOffset: 15

- name: Grassland
  sprites:
  - name: Grassland
    xOffset: 0
    yOffset: 0
`
	_, cleanupUIFile := testutils.CreateFileInDir(t, dirPath, "ui.yml", uiContent)
	defer cleanupUIFile()

	// Now you can test LoadRuleset with the path to the temporary directory
	_, err := LoadRuleset(dirPath)
	if err != nil {
		t.Errorf("LoadRuleset failed: %v", err)
	}

	// Additional assertions to verify the behavior of LoadRuleset
}

func TestRuleset_FindTileType(t *testing.T) {

	// Setup a test gameRuleset
	ruleset := Ruleset{
		EnabledTiles: []string{"Grassland", "Forest", "Mountain"},
	}

	tests := []struct {
		name      string
		tileName  string
		want      game.TileType
		wantFound bool
	}{
		{
			name:      "TileExists",
			tileName:  "Forest",
			want:      1, // Assuming the index (TileType) of "Forest" is 1
			wantFound: true,
		},
		{
			name:      "TileDoesNotExist",
			tileName:  "Desert",
			want:      game.UnknownTileType, // Assuming UnknownTileType represents not found
			wantFound: false,
		},
		{
			name:      "FirstTile",
			tileName:  "Grassland",
			want:      0, // Assuming the index (TileType) of "Grassland" is 0
			wantFound: true,
		},
		{
			name:      "LastTile",
			tileName:  "Mountain",
			want:      2, // Assuming the index (TileType) of "Mountain" is 2
			wantFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := ruleset.FindTileType(tt.tileName)
			if got != tt.want || found != tt.wantFound {
				t.Errorf("FindTileType(%v) got = %v, found = %v, want %v, wantFound %v", tt.tileName, got, found, tt.want, tt.wantFound)
			}
		})
	}

}

func TestRuleset_FindModifierType(t *testing.T) {
	// Setup a test gameRuleset with known modifiers
	ruleset := Ruleset{
		EnabledModifiers: []string{"River", "Road", "MountainPass"},
	}

	tests := []struct {
		name         string
		modifierName string
		want         game.ModifierType
		wantFound    bool
	}{
		{
			name:         "ModifierExists",
			modifierName: "River",
			want:         0, // Assuming the index (ModifierType) of "River" is 0
			wantFound:    true,
		},
		{
			name:         "ModifierDoesNotExist",
			modifierName: "Bridge",
			want:         game.UnknownModifierType, // Assuming UnknownModifierType represents not found
			wantFound:    false,
		},
		{
			name:         "LastModifier",
			modifierName: "MountainPass",
			want:         2, // Assuming the index (ModifierType) of "MountainPass" is 2
			wantFound:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := ruleset.FindModifierType(tt.modifierName)
			if got != tt.want || found != tt.wantFound {
				t.Errorf("FindModifierType(%v) got = %v, found = %v, want %v, wantFound %v", tt.modifierName, got, found, tt.want, tt.wantFound)
			}
		})
	}
}
