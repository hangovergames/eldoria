// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package testutils

import (
	"log"
	"os"
	"testing"
)

func CreateTempYAMLFile(t *testing.T, content string) (filename string, cleanupFunc func()) {

	t.Helper()

	tempFile, err := os.CreateTemp("", "example.*.yaml")
	if err != nil {
		log.Fatalf("Failed to create temp file: %v", err)
	}

	if _, err := tempFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
		tempFile.Close()
	}

	if err := tempFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tempFile.Name(), func() {
		os.Remove(tempFile.Name())
	}
}
