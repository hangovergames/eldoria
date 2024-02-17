// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package testutils

import (
	"os"
	"path/filepath"
	"testing"
)

func CreateTempDir(t *testing.T) (dirPath string, cleanupFunc func()) {
	t.Helper()

	dirPath, err := os.MkdirTemp("", "tempDir")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}

	cleanupFunc = func() {
		os.RemoveAll(dirPath) // Cleanup the directory after test
	}

	return dirPath, cleanupFunc
}

func CreateFileInDir(t *testing.T, dirPath, fileName, content string) (filePath string, cleanupFunc func()) {
	t.Helper()

	filePath = filepath.Join(dirPath, fileName)
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write to file %s: %v", filePath, err)
	}

	cleanupFunc = func() {
		os.Remove(filePath) // Cleanup the file after test
	}

	return filePath, cleanupFunc
}
