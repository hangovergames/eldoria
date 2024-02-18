// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package main

import (
	"fmt"
	"github.com/hangovergames/eldoria/internal/client/fontutils"
	"github.com/hangovergames/eldoria/internal/client/screens/loginscreen"
	"github.com/hangovergames/eldoria/internal/client/ui"
	"github.com/hangovergames/eldoria/internal/common/assets"
	"github.com/hangovergames/eldoria/internal/common/assets/fonts"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/freeciv/data/trident"

	"github.com/hangovergames/eldoria/internal/client/gameui"
	"github.com/hangovergames/eldoria/internal/client/imageutils"
	"github.com/hangovergames/eldoria/internal/client/spriteutils"
	"github.com/hangovergames/eldoria/internal/client/uimap"
	"github.com/hangovergames/eldoria/internal/common/apiClient"

	"github.com/hangovergames/eldoria/internal/client/screens/mapScreen"
)

var (
	imageManager *imageutils.ImageManager
	gameUI       *gameui.GameUI
)

func init() {

}

func main() {

	fontManager := fontutils.NewFontManager()
	fontManager.RegisterFontBytes(ui.IndieFlowerFontName, fonts.IndieFlower_Regular_ttf)
	fontManager.RegisterFontBytes(ui.MagicSchoolFontName, fonts.MagicSchoolTwo_ttf)
	fontManager.RegisterFontBytes(ui.PlayfairDisplayFontName, fonts.PlayfairDisplay_Regular_ttf)
	fontManager.RegisterFontBytes(ui.PatrickHandFontName, fonts.PatrickHand_Regular_ttf)
	fontManager.RegisterFontBytes(ui.OrbitronFontName, fonts.Orbitron_Regular_ttf)
	fontManager.RegisterFontBytes(ui.MerriweatherFontName, fonts.Merriweather_Regular_ttf)

	imageManager = imageutils.NewImageManager()
	imageManager.RegisterImageBytes("freeciv/data/trident/tiles.png", images.Tiles_png)
	imageManager.RegisterImageBytes("freeciv/data/trident/units.png", images.Units_png)
	imageManager.RegisterImageBytes("freeciv/data/trident/extra_units.png", images.ExtraUnits_png)
	imageManager.RegisterImageBytes("freeciv/data/trident/fog.png", images.Fog_png)
	imageManager.RegisterImageBytes("freeciv/data/trident/select.png", images.Select_png)
	imageManager.RegisterImageBytes("freeciv/data/trident/roads.png", images.Roads_png)
	imageManager.RegisterImageBytes("freeciv/data/trident/grid.png", images.Grid_png)
	imageManager.RegisterImageBytes("freeciv/data/trident/explosions.png", images.Explosions_png)
	imageManager.RegisterImageBytes("freeciv/data/trident/earth.png", images.Earth_png)
	imageManager.RegisterImageBytes("freeciv/data/trident/cities.png", images.Cities_png)
	imageManager.RegisterImageBytes(ui.LoginScreenBackgroundImage, assets.EldoriaLoginScreen_png)

	client := apiclient.NewAPIClient("http://localhost:8080")

	uiConfig, err := client.FetchUIConfigDTO()
	if err != nil {
		fmt.Printf("Error fetching UI config: %v\n", err)
		return
	}

	spriteManager := spriteutils.NewSpriteManager(imageManager)
	spriteManager.LoadSpriteSheetDTOs(uiConfig.SpriteSheets)
	spriteManager.LoadSpriteConfigDTOs(uiConfig.SpriteConfigs)

	tileMap := uimap.NewTileGrid(spriteManager, 10, 10)
	tileMap.LoadTileConfigDTOs(uiConfig.TileConfigs)

	// Draw tiles
	tileMap.SetTile(5, 5, "Grassland", "Warrior")

	gameUI = gameui.NewGameUI(ui.InitialScreenWidth, ui.InitialScreenHeight, tileMap)

	gameUI.RegisterScreen("Login", func() ui.IScreen {
		return loginscreen.NewLoginScreen(
			gameUI,
			imageManager,
			fontManager,
			ui.LoginScreenBackgroundImage,
			ui.LoginScreenFontName,
			ui.LoginScreenFontSize,
			ui.LoginScreenFontDPI,
			ui.MinLoginNameLength,
			ui.MaxLoginNameLength,
			ui.AllowedLoginCharacters,
		)
	})

	gameUI.RegisterScreen("Map", func() ui.IScreen {
		return mapscreen.NewMapScreen(gameUI)
	})

	gameUI.SetCurrentScreen("Login")

	ebiten.SetWindowSize(ui.InitialScreenWidth, ui.InitialScreenHeight)
	ebiten.SetWindowTitle("Eldoria (Client)")
	if err := ebiten.RunGame(gameUI); err != nil {
		log.Fatal(err)
	}

}
