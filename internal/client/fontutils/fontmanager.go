// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package fontutils

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"log"
	"os"
)

// FontManager manages the storage and retrieval of truetype.Fonts by name.
type FontManager struct {
	fonts map[string]*truetype.Font
}

// NewFontManager creates a new instance of FontManager.
func NewFontManager() *FontManager {
	return &FontManager{
		fonts: make(map[string]*truetype.Font),
	}
}

// RegisterFont associates a *truetype.Font with a name.
func (fm *FontManager) RegisterFont(name string, font *truetype.Font) {
	if font == nil {
		log.Fatal("RegisterFont called with null font")
	}
	fm.fonts[name] = font
}

// GetFont retrieves a *truetype.Font by name.
func (fm *FontManager) GetFont(name string) *truetype.Font {
	fnt, exists := fm.fonts[name]
	if exists {
		return fnt
	}
	return nil
}

// LoadFont loads a font from a file and registers it under the given name.
func (fm *FontManager) LoadFont(name, filePath string) error {
	fontBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read font file: %v", err)
		return err
	}

	fnt, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Printf("Failed to parse font: %v", err)
		return err
	}

	fm.RegisterFont(name, fnt)
	return nil
}

// GetFace returns a font.Face for a registered font with the specified size and DPI.
func (fm *FontManager) GetFace(name string, size float64, dpi float64) font.Face {
	fnt := fm.GetFont(name)
	if fnt == nil {
		log.Printf("Font not found: %s", name)
		return nil
	}
	// truetype.Options allows you to specify the size and DPI for the font face.
	opts := &truetype.Options{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	}
	return truetype.NewFace(fnt, opts)
}

// RegisterFontBytes registers a font from a byte slice.
func (fm *FontManager) RegisterFontBytes(name string, fontBytes []byte) {
	fnt, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("Failed to parse font bytes for %s: %v", name, err)
	}
	fm.fonts[name] = fnt
}
