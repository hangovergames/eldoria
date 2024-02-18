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
	TilesPngImg, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}

	// Decode an image from the image file's byte slice.
	UnitsPngImg, _, err := image.Decode(bytes.NewReader(images.Units_png))
	if err != nil {
		log.Fatal(err)
	}

	// Decode an image from the image file's byte slice.
	ExtraUnitsPngImg, _, err := image.Decode(bytes.NewReader(images.ExtraUnits_png))
	if err != nil {
		log.Fatal(err)
	}

	// Decode an image from the image file's byte slice.
	FogPngImg, _, err := image.Decode(bytes.NewReader(images.Fog_png))
	if err != nil {
		log.Fatal(err)
	}

	// Decode an image from the image file's byte slice.
	SelectPngImg, _, err := image.Decode(bytes.NewReader(images.Select_png))
	if err != nil {
		log.Fatal(err)
	}

	// Decode an image from the image file's byte slice.
	RoadsPngImg, _, err := image.Decode(bytes.NewReader(images.Roads_png))
	if err != nil {
		log.Fatal(err)
	}

	// Decode an image from the image file's byte slice.
	GridPngImg, _, err := image.Decode(bytes.NewReader(images.Grid_png))
	if err != nil {
		log.Fatal(err)
	}

	// Decode an image from the image file's byte slice.
	ExplosionsPngImg, _, err := image.Decode(bytes.NewReader(images.Explosions_png))
	if err != nil {
		log.Fatal(err)
	}

	// Decode an image from the image file's byte slice.
	EarthPngImg, _, err := image.Decode(bytes.NewReader(images.Earth_png))
	if err != nil {
		log.Fatal(err)
	}

	// Decode an image from the image file's byte slice.
	CitiesPngImg, _, err := image.Decode(bytes.NewReader(images.Cities_png))
	if err != nil {
		log.Fatal(err)
	}

	imageManager = imageutils.NewImageManager()
	imageManager.RegisterImage("freeciv/data/trident/tiles.png", ebiten.NewImageFromImage(TilesPngImg))
	imageManager.RegisterImage("freeciv/data/trident/units.png", ebiten.NewImageFromImage(UnitsPngImg))
	imageManager.RegisterImage("freeciv/data/trident/extra_units.png", ebiten.NewImageFromImage(ExtraUnitsPngImg))
	imageManager.RegisterImage("freeciv/data/trident/fog.png", ebiten.NewImageFromImage(FogPngImg))
	imageManager.RegisterImage("freeciv/data/trident/select.png", ebiten.NewImageFromImage(SelectPngImg))
	imageManager.RegisterImage("freeciv/data/trident/roads.png", ebiten.NewImageFromImage(RoadsPngImg))
	imageManager.RegisterImage("freeciv/data/trident/grid.png", ebiten.NewImageFromImage(GridPngImg))
	imageManager.RegisterImage("freeciv/data/trident/explosions.png", ebiten.NewImageFromImage(ExplosionsPngImg))
	imageManager.RegisterImage("freeciv/data/trident/earth.png", ebiten.NewImageFromImage(EarthPngImg))
	imageManager.RegisterImage("freeciv/data/trident/cities.png", ebiten.NewImageFromImage(CitiesPngImg))

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
	tileMap.SetTile(5, 5, "Grassland", "Warrior")

	gameUI = gameui.NewGameUI(screenWidth, screenHeight, tileMap)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Eldoria (Client)")
	if err := ebiten.RunGame(gameUI); err != nil {
		log.Fatal(err)
	}

}
