// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package fontutils

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"testing"
)

func TestFontManager(t *testing.T) {
	fm := NewFontManager()

	// Test RegisterFont and GetFont
	fontName := "testFont"
	testFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		t.Fatalf("Failed to parse test font: %v", err)
	}
	fm.RegisterFont(fontName, testFont)
	retrievedFont := fm.GetFont(fontName)
	if retrievedFont != testFont {
		t.Errorf("Retrieved font does not match registered font")
	}

	// Test RegisterFontBytes
	fontBytesName := "testFontBytes"
	fm.RegisterFontBytes(fontBytesName, goregular.TTF)
	retrievedFontBytes := fm.GetFont(fontBytesName)
	if retrievedFontBytes == nil {
		t.Errorf("Failed to retrieve font registered from bytes")
	}

	// Test GetFace
	size := 12.0 // Example font size
	dpi := 72.0  // Example DPI
	face := fm.GetFace(fontName, size, dpi)
	if face == nil {
		t.Errorf("Failed to create font face")
	}
}
