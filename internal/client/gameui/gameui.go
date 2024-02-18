// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hangovergames/eldoria/internal/client/ui"
	"log"
)

type GameUI struct {
	ScreenWidth, ScreenHeight int
	Map                       ui.ITileGrid
	CurrentScreen             ui.IScreen
	screens                   map[string]func() ui.IScreen
}

// NewGameUI creates a new instance of GameUI.
func NewGameUI(
	screenWidth, screenHeight int,
	grid ui.ITileGrid,
) *GameUI {
	return &GameUI{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
		Map:          grid,
		screens:      make(map[string]func() ui.IScreen),
	}
}

// RegisterScreen registers a screen constructor function under a given name.
func (g *GameUI) RegisterScreen(name string, constructor func() ui.IScreen) {
	g.screens[name] = constructor
}

// GetScreen retrieves a screen by name, instantiating it if necessary.
func (g *GameUI) GetScreen(name string) ui.IScreen {
	if screenConstructor, ok := g.screens[name]; ok {
		// Instantiate the screen if it's not already active or if you want a new instance every time.
		return screenConstructor()
	}
	// Handle the case where no screen is found for the given name, possibly returning nil or an error.
	return nil
}

func (g *GameUI) Update() error {
	if g.CurrentScreen != nil {
		return g.CurrentScreen.Update()
	}
	return nil
}

func (g *GameUI) Draw(screen *ebiten.Image) {
	if g.CurrentScreen != nil {
		g.CurrentScreen.Draw(screen)
	}
}

func (g *GameUI) Layout(outsideWidth, outsideHeight int) (int, int) {
	if g.CurrentScreen != nil {
		return g.CurrentScreen.Layout(outsideWidth, outsideHeight)
	}
	return g.GetScreenWidth(), g.GetScreenHeight()
}

func (g *GameUI) GetMap() ui.ITileGrid {
	if g.CurrentScreen != nil {
		return g.Map
	}
	return nil
}

func (g *GameUI) GetScreenWidth() int {
	return g.ScreenWidth
}

func (g *GameUI) GetScreenHeight() int {
	return g.ScreenHeight
}

func (g *GameUI) GetCurrentScreen() ui.IScreen {
	return g.CurrentScreen
}

func (g *GameUI) SetCurrentScreen(name string) {
	screen := g.GetScreen(name)
	if screen != nil {
		g.CurrentScreen = screen
	} else {
		log.Printf("Could not create a screen named %s", name)
	}
}
