// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package main

import (
	"bytes"
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/freeciv/data/trident"

	"github.com/hangovergames/eldoria/internal/client/gameui"
	"github.com/hangovergames/eldoria/internal/client/imageutils"
	"github.com/hangovergames/eldoria/internal/client/spriteutils"
	"github.com/hangovergames/eldoria/internal/client/ui/uiMap"
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

	uiConfig := dtos.LoadUIConfigDTO("./examples/ui-config-dto.json")

	imageManager.RegisterImage("freeciv/data/trident/tiles.png", ebiten.NewImageFromImage(img))

	spriteManager := spriteutils.NewSpriteManager(imageManager)

	spriteManager.LoadSpriteSheetDTOs(uiConfig.SpriteSheets)

	spriteManager.MapSpriteName("ShallowOcean", "OceanTiles", 0)
	spriteManager.MapSpriteName("DeepOcean", "OceanTiles", 10)
	spriteManager.MapSpriteName("Grassland", "Tiles", 2)

	tileMap := uiMap.NewTileGrid(spriteManager, 10, 10)
	tileMap.LoadTileConfigDTOs(uiConfig.TileConfigs)

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
