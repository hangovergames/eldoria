// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameui

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hangovergames/eldoria/internal/client/ui/uiMap"
)

type GameUI struct {
	ScreenWidth, ScreenHeight int
	Map                       *uiMap.TileGrid
}

// NewGameUI creates a new instance of GameUI.
func NewGameUI(
	screenWidth, screenHeight int,
	grid *uiMap.TileGrid,
) *GameUI {
	return &GameUI{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
		Map:          grid,
	}
}

func (g *GameUI) Update() error {
	return nil
}

func (g *GameUI) Draw(screen *ebiten.Image) {

	// Render the map onto the screen
	if g.Map != nil {
		g.Map.Draw(screen, 30, 30)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))

}

func (g *GameUI) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
