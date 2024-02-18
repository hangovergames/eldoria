// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameruleset

import (
	"github.com/hangovergames/eldoria/internal/common/testutils"
	"testing"
)

func TestLoadResourcesFile(t *testing.T) {
	// Success case with valid YAML content
	t.Run("ValidYAML", func(t *testing.T) {
		content := `
resources:
  Gold:
    summary: "Precious metal used for trade and crafting."
  Wood:
    summary: "Basic building material."
`
		filename, cleanup := testutils.CreateTempYAMLFile(t, content)
		defer cleanup()

		got, err := LoadResourcesFile(filename)
		if err != nil {
			t.Fatalf("LoadResourcesFile() error = %v, wantErr false", err)
		}
		if len(got.Resources) != 2 {
			t.Errorf("Expected 2 resources, got %d", len(got.Resources))
		}
	})

	// Error case with invalid YAML content
	t.Run("InvalidYAML", func(t *testing.T) {
		content := `
resources:
  Gold:
    summary: "Precious metal used for trade and crafting."
  Wood: "Basic building material." # Missing summary key
`
		filename, cleanup := testutils.CreateTempYAMLFile(t, content)
		defer cleanup()

		_, err := LoadResourcesFile(filename)
		if err == nil {
			t.Fatal("Expected error for invalid YAML, got nil")
		}
	})

	// Error case with nonexistent file
	t.Run("NonexistentFile", func(t *testing.T) {
		_, err := LoadResourcesFile("nonexistent.yml")
		if err == nil {
			t.Fatal("Expected error for nonexistent file, got nil")
		}
	})
}
