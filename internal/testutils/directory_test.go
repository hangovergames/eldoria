// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package testutils

import (
	"os"
	"testing"
)

func TestCreateTempDir(t *testing.T) {
	dirPath, cleanupFunc := CreateTempDir(t)
	defer cleanupFunc()

	// Verify the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		t.Errorf("Temporary directory was not created")
	}

	// Call cleanup and verify the directory is removed
	cleanupFunc()
	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		t.Errorf("Temporary directory was not removed by cleanup")
	}
}

func TestCreateFileInDir(t *testing.T) {
	// First, create a temporary directory to hold the file
	dirPath, dirCleanupFunc := CreateTempDir(t)
	defer dirCleanupFunc()

	// Define the file name and content
	fileName := "testFile.txt"
	fileContent := "Hello, world!"

	// Create the file within the temporary directory
	filePath, fileCleanupFunc := CreateFileInDir(t, dirPath, fileName, fileContent)
	defer fileCleanupFunc()

	// Read back the content to verify it matches
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read the file: %v", err)
	}
	if string(content) != fileContent {
		t.Errorf("File content does not match. Expected %s, got %s", fileContent, string(content))
	}

	// Verify the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("File was not created")
	}

	// Call cleanup and verify the file is removed
	fileCleanupFunc()
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Errorf("File was not removed by cleanup")
	}
}
