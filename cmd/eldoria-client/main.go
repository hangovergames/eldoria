// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package main

import (
	"bytes"
	"fmt"
	"github.com/hangovergames/eldoria/internal/common/apiClient"
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

}

func main() {

	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}

	imageManager = imageutils.NewImageManager()
	imageManager.RegisterImage("freeciv/data/trident/tiles.png", ebiten.NewImageFromImage(img))

	client := apiClient.NewAPIClient("http://localhost:8080")

	uiConfig, err := client.FetchUIConfigDTO()
	if err != nil {
		fmt.Printf("Error fetching UI config: %v\n", err)
		return
	}

	spriteManager := spriteutils.NewSpriteManager(imageManager)
	spriteManager.LoadSpriteSheetDTOs(uiConfig.SpriteSheets)
	spriteManager.LoadSpriteConfigDTOs(uiConfig.SpriteConfigs)

	tileMap := uiMap.NewTileGrid(spriteManager, 10, 10)
	tileMap.LoadTileConfigDTOs(uiConfig.TileConfigs)

	// Draw tiles
	tileMap.SetTile(5, 5, "Grassland")

	gameUI = gameui.NewGameUI(screenWidth, screenHeight, tileMap)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Eldoria (Client)")
	if err := ebiten.RunGame(gameUI); err != nil {
		log.Fatal(err)
	}

}
