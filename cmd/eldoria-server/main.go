// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package main

import (
	"fmt"
	"log"

	"github.com/hangovergames/eldoria/internal/server/game"
	"github.com/hangovergames/eldoria/internal/server/gameruleset"
	"github.com/hangovergames/eldoria/internal/server/gameserver"
	"github.com/hangovergames/eldoria/internal/server/gamestate"

	"github.com/hangovergames/eldoria/internal/server/gamemap"
)

const PORT = "8080"

func main() {

	ruleset, err := gameruleset.LoadRuleset("./ruleset/default")
	if err != nil {
		log.Fatalf("failed to load ruleset: %v", err)
	}

	// Start the HTTP gameServer.
	address := fmt.Sprintf(":%s", PORT)

	gameMap := gamemap.NewGameMap(10, 10, gamemap.NewTile(0, []game.ModifierType{}, 0))

	state := gamestate.NewGameState(gameMap)

	server := gameserver.NewServer(address, &ruleset, state)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
