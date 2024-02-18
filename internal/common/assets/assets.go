// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package assets

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

var (

	//go:embed eldoria-login-screen.png
	EldoriaLoginScreen_png []byte

	WhiteImage = ebiten.NewImage(3, 3)

	// whiteSubImage is an internal sub image of whiteImage.
	// Use whiteSubImage at DrawTriangles instead of whiteImage in order to avoid bleeding edges.
	WhiteSubImage = WhiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	WhiteImage.Fill(color.White)
}
