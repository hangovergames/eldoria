// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package testutils

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestCreateTempYAMLFile verifies that the CreateTempYAMLFile function creates a file with the expected content
// and that the cleanup function removes the file.
func TestCreateTempYAMLFile(t *testing.T) {
	// Define the content to write to the temporary YAML file.
	testContent := "key: value"

	// Call the function to create a temporary YAML file.
	filename, cleanup := CreateTempYAMLFile(t, testContent)
	defer cleanup() // Ensure cleanup runs after the test to remove the temporary file.

	// Read back the content from the temporary file to verify it matches what was expected.
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read back the temporary file: %v", err)
	}

	if string(content) != testContent {
		t.Errorf("Content of the temp file does not match expected content. Got: %s, Want: %s", string(content), testContent)
	}

	// Verify that the file exists before cleanup.
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("Temporary file does not exist: %s", filename)
	}

	// Call the cleanup function to remove the temporary file.
	cleanup()

	// Verify that the file has been removed after cleanup.
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		t.Errorf("Temporary file was not cleaned up: %s", filename)
	}
}
