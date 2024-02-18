// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mapscreen

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hangovergames/eldoria/internal/client/ui"
)

type MapScreen struct {
	Game ui.IGame
}

func NewMapScreen(game ui.IGame) *MapScreen {
	return &MapScreen{
		Game: game,
	}
}

func (m *MapScreen) Update() error {
	// Handle updates for the menu screen
	return nil
}

func (m *MapScreen) Draw(screen *ebiten.Image) {

	myMap := m.Game.GetMap()
	if myMap != nil {
		myMap.Draw(screen, 30, 30)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))

}

func (m *MapScreen) Layout(outsideWidth, outsideHeight int) (int, int) {
	return m.Game.GetScreenWidth(), m.Game.GetScreenHeight()
}
