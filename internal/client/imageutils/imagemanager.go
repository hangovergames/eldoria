// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package imageutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
	"os"
)

// ImageManager manages the storage and retrieval of ebiten.Images by name.
type ImageManager struct {
	images map[string]*ebiten.Image
}

// NewImageManager creates a new instance of ImageManager.
func NewImageManager() *ImageManager {
	return &ImageManager{
		images: make(map[string]*ebiten.Image),
	}
}

// RegisterImage associates an *ebiten.Image with a name.
func (im *ImageManager) RegisterImage(name string, img *ebiten.Image) {
	if img == nil {
		log.Fatal("RegisterImage called with null image")
	}
	im.images[name] = img
}

// GetImage retrieves an *ebiten.Image by name.
func (im *ImageManager) GetImage(name string) *ebiten.Image {
	img, exists := im.images[name]
	if exists {
		return img
	}
	return nil
}

// LoadImage loads an image from a file and registers it under the given name.
func (im *ImageManager) LoadImage(name, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open image file: %v", err)
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Printf("Failed to decode image file: %v", err)
		return err
	}

	ebitenImage := ebiten.NewImageFromImage(img)
	im.RegisterImage(name, ebitenImage)
	return nil
}
