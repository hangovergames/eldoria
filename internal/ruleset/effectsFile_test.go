// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package ruleset

import (
	"github.com/hangovergames/eldoria/internal/testutils"
	"reflect"
	"testing"
)

func TestLoadEffectsFile(t *testing.T) {

	// Define a test case for loading a valid effects YAML file
	t.Run("ValidEffectsFile", func(t *testing.T) {
		content := `
effects:
  ProvidesFreshWater:
    summary: "Crucial for life and agriculture."
    resources:
      FreshWater:
        restorationRate: {min: 10, max: 10}
`
		filename, cleanup := testutils.CreateTempYAMLFile(t, content)
		defer cleanup()

		got, err := LoadEffectsFile(filename)
		if err != nil {
			t.Errorf("LoadEffectsFile() error = %v, wantErr false", err)
			return
		}

		want := EffectsFile{
			Effects: map[string]EffectConfig{
				"ProvidesFreshWater": {
					Summary: "Crucial for life and agriculture.",
					Resources: map[string]TileResourceConfig{
						"FreshWater": {
							RestorationRate: Rate{Min: 10, Max: 10},
						},
					},
				},
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("LoadEffectsFile() got = %v, want %v", got, want)
		}
	})

	// Add more test cases as needed, such as "InvalidYAMLFormat", "FileDoesNotExist", etc.

}
