// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package game

import (
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"net/url"
)

type PlayerID uint
type TileType uint
type ModifierType uint
type TileEffect uint

// HasTileEffect Function to check if a tile effect includes a specific effect
func (t TileEffect) HasTileEffect(effect TileEffect) bool {
	return t&effect != 0
}

type IResponse interface {
	Send(statusCode int, data interface{})
	SendError(statusCode int, error string)
	SendMethodNotSupportedError()
}

// IServer defines the methods available from the Server
// that are needed by the HTTP handlers.
type IServer interface {
	Start() error
	SetupRoutes()
	GetAddress() string
	GetRuleset() IRuleset
	GetState() IGameState
}

// IRuleset defines the methods that the Ruleset needs to expose to external consumers.
type IRuleset interface {
	FindTileType(tileName string) (TileType, bool)
	FindModifierType(tileName string) (ModifierType, bool)
	CreateTileFromName(tileName string) (ITile, error)
	IsEffectEnabled(effectStr string) bool
	ConvertAndCombineTileEffects(effectStrings []string) (TileEffect, error)
	GetUI() dtos.UIConfigDTO
}

// ITile defines the interface for operations on a Tile.
type ITile interface {

	// GetType returns the type of the tile.
	GetType() TileType

	// GetModifiers returns the slice of ModifierType applied to the tile.
	GetModifiers() []ModifierType

	// GetEffects returns the TileEffect applied to the tile.
	GetEffects() TileEffect

	// Clone creates and returns a deep copy of the tile.
	Clone() ITile
}

// IGameMap defines the interface for interacting with a game map.
type IGameMap interface {
	GetWidth() int
	GetHeight() int

	// GetTile returns the tile at the given coordinates.
	GetTile(x, y int) (ITile, error)

	// SetTile sets a tile at the given coordinates.
	SetTile(x, y int, newTile ITile) error

	// GetTilesInArea returns a 2D slice of tiles within the specified rectangular area.
	GetTilesInArea(x, y, x2, y2 int) ([][]ITile, error)
}

// IPlayer defines the interface for interacting with a player.
type IPlayer interface {
	// GetID returns the player's ID.
	GetID() PlayerID

	// GetName returns the player's name.
	GetName() string

	// SetName sets the player's name.
	SetName(name string)

	// GetResources returns the player's resources.
	GetResources() map[string]int

	// AddResource adds a specified amount of a resource to the player's inventory.
	AddResource(resource string, amount int)

	// SubtractResource subtracts a specified amount of a resource from the player's inventory.
	SubtractResource(resource string, amount int)
}

// IGameState defines the interface for interacting with the game state.
type IGameState interface {

	// GetMap returns the current game map.
	GetMap() IGameMap

	// GetPlayers returns a list of all players in the game state.
	GetPlayers() []IPlayer

	// FindPlayer returns a player if it exists
	FindPlayer(name string) (IPlayer, error)

	// AddPlayer adds a new player to the game state.
	//AddPlayer(player gamePlayer.Player)

	// RemovePlayer removes a player from the game state by their identifier.
	//RemovePlayer(name string)

	// UpdatePlayer updates a player's information in the game state.
	//UpdatePlayer(player gamePlayer.Player)

}

type IRequest interface {
	IsMethodGet() bool
	GetURL() *url.URL
	GetVars() map[string]string
}

// RequestHandlerFunc defines the type for handlers in this API.
type RequestHandlerFunc func(IResponse, IRequest, IServer)
