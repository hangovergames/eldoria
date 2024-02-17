// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package main

import (
	"bytes"
	"github.com/hangovergames/eldoria/internal/client/gameui"
	"github.com/hangovergames/eldoria/internal/client/imageutils"
	"github.com/hangovergames/eldoria/internal/client/spriteutils"
	"github.com/hangovergames/eldoria/internal/client/ui/uiMap"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/freeciv/data/trident"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var (
	imageManager *imageutils.ImageManager
	gameUI       *gameui.GameUI
)

func init() {

	imageManager = imageutils.NewImageManager()

	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}

	imageManager.RegisterImage("freeciv/data/trident/tiles.png", ebiten.NewImageFromImage(img))

	tilesSheet := spriteutils.NewSpriteSheet(
		imageManager.GetImage("freeciv/data/trident/tiles.png"),
		30,
		30,
		20,
		0,
		0,
	)

	oceanTilesSheet := spriteutils.NewSpriteSheet(
		imageManager.GetImage("freeciv/data/trident/tiles.png"),
		15,
		15,
		32,
		0,
		210,
	)

	spriteManager := spriteutils.NewSpriteManager()

	spriteManager.RegisterSpriteSheet("Tiles", tilesSheet)
	spriteManager.RegisterSpriteSheet("OceanTiles", oceanTilesSheet)

	spriteManager.MapSpriteName("ShallowOcean", "OceanTiles", 0)
	spriteManager.MapSpriteName("DeepOcean", "OceanTiles", 10)
	spriteManager.MapSpriteName("Grassland", "Tiles", 2)

	tileMap := uiMap.NewTileGrid(spriteManager, 10, 10)

	tileMap.DefineTileConfig("DeepOcean", "DeepOcean", 0, 0)
	tileMap.DefineTileConfig("DeepOcean", "DeepOcean", 15, 0)
	tileMap.DefineTileConfig("DeepOcean", "DeepOcean", 0, 15)
	tileMap.DefineTileConfig("DeepOcean", "DeepOcean", 15, 15)

	tileMap.DefineTileConfig("Grassland", "Grassland", 0, 0)

	tileMap.SetTile(5, 5, "Grassland")

	gameUI = gameui.NewGameUI(screenWidth, screenHeight, tileMap)

}

func main() {

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Eldoria (Client)")
	if err := ebiten.RunGame(gameUI); err != nil {
		log.Fatal(err)
	}

}
