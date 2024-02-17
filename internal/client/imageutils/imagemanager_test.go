// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package imageutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestImageManager_RegisterAndGetImage(t *testing.T) {

	manager := NewImageManager()

	// Mock or use a minimal ebiten.Image for testing; here, we're not focusing on the image content.
	mockImage := ebiten.NewImage(1, 1) // Creating a minimal image; in real tests, consider mocking.

	name := "testImage"
	manager.RegisterImage(name, mockImage)

	retrievedImage := manager.GetImage(name)
	if retrievedImage == nil {
		t.Errorf("GetImage() failed to retrieve the registered image")
	}

	// Test retrieving a non-existent image
	if image := manager.GetImage("nonExistent"); image != nil {
		t.Errorf("GetImage() should return false for non-existent images")
	}

}
