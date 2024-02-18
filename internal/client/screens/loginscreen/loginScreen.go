// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package loginscreen

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hangovergames/eldoria/internal/client/ui"
	"image/color"
	"strings"
)

type LoginScreen struct {
	Game                ui.IGame
	ImageManager        ui.IImageManager
	BackgroundImageName string
	ShowPopup           bool
	UsernameField       TextField
	FontManager         ui.IFontManager
	FontName            string
	FontSize            float64
	FontDPI             float64
	MinLoginLength      int
	MaxLoginLength      int
}

func NewLoginScreen(
	game ui.IGame,
	imageManager ui.IImageManager,
	fontManager ui.IFontManager,
	backgroundImageName string,
	fontName string,
	fontSize float64,
	fontDPI float64,
	minLoginLength int,
	maxLoginLength int,
	allowedLoginCharacters string,
) *LoginScreen {

	ls := &LoginScreen{
		Game:                game,
		ImageManager:        imageManager,
		FontManager:         fontManager,
		BackgroundImageName: backgroundImageName,
		ShowPopup:           false,
		UsernameField: TextField{
			MinLength:           minLoginLength,
			MaxLength:           maxLoginLength,
			X:                   0,
			Y:                   0,
			Width:               200,
			Height:              30,
			FontManager:         fontManager,
			FontName:            fontName,
			PlaceholderFontName: fontName,
			Placeholder:         "Login name",
			AllowedChars:        allowedLoginCharacters,
		},
		FontName: fontName,
		FontSize: fontSize,
		FontDPI:  fontDPI,
	}

	ls.UsernameField.SetFont(fontName, fontSize, fontDPI, color.Black)
	ls.UsernameField.SetPlaceholderFont(fontName, fontSize, fontDPI, color.RGBA{R: 150, G: 150, B: 150, A: 255})

	ls.UsernameField.OnEnter = func() { ls.onUsernameEnter() }

	return ls
}

func (m *LoginScreen) Update() error {

	// Check if the left mouse button was just pressed
	if !m.ShowPopup {

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			m.ShowPopup = true
			m.UsernameField.IsActive = true
		}

	} else {
		m.UsernameField.Update()
	}

	// Handle updates for the menu screen
	return nil
}

func (m *LoginScreen) Draw(screen *ebiten.Image) {

	bgImage := m.ImageManager.GetImage(m.BackgroundImageName)
	if bgImage != nil {

		screenWidth, screenHeight := m.Game.GetScreenWidth(), m.Game.GetScreenHeight()
		imgWidth, imgHeight := bgImage.Bounds().Dx(), bgImage.Bounds().Dy()

		// Calculate the scaling factor to match the screen height while preserving the aspect ratio
		scale := float64(screenHeight) / float64(imgHeight)

		// Calculate new width after scaling to check if it exceeds the screen width
		scaledWidth := float64(imgWidth) * scale

		// Calculate horizontal offset to center the image
		offsetX := (float64(screenWidth) - scaledWidth) / 2

		// Create an ebiten.DrawImageOptions object to apply transformations
		opts := &ebiten.DrawImageOptions{}
		// Scale the image
		opts.GeoM.Scale(scale, scale)
		// Apply the calculated offset (if the scaled image width is greater than the screen width)
		opts.GeoM.Translate(offsetX, 0)

		// Draw the background image with the options
		screen.DrawImage(bgImage, opts)

	} else {
		fmt.Printf("Background image not found: %s", m.BackgroundImageName)
	}

	if m.ShowPopup {

		// Render the popup
		popupWidth, popupHeight := 200, 100 // Set your desired dimensions
		popupX, popupY := (float64(m.Game.GetScreenWidth())-float64(popupWidth))/2, (float64(m.Game.GetScreenHeight())-float64(popupHeight))/2

		// For example, draw a rectangle as the popup background
		//popupImage := ebiten.NewImage(popupWidth, popupHeight)
		//popupImage.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 255}) // White background

		//opts := &ebiten.DrawImageOptions{}
		//opts.GeoM.Translate(popupX, popupY)
		//screen.DrawImage(popupImage, opts)

		m.UsernameField.X = int(popupX) + 10
		m.UsernameField.Y = int(popupY) + 10
		m.UsernameField.Draw(screen)

	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (m *LoginScreen) Layout(outsideWidth, outsideHeight int) (int, int) {
	return m.Game.GetScreenWidth(), m.Game.GetScreenHeight()
}

// onUsernameEnter defines what happens when the Enter key is pressed in the username field
func (m *LoginScreen) onUsernameEnter() {

	username := strings.ToLower(m.UsernameField.Text)

	// Implement the action, for example, logging in
	fmt.Println("Enter pressed, username:", username)

	// You can call a login function or whatever you need to do here

}
