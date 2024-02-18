// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package main

import (
	"fmt"
	"log"

	"github.com/hangovergames/eldoria/internal/server/gameRuleset"
	"github.com/hangovergames/eldoria/internal/server/gameServer"
)

const PORT = "8080"

func main() {

	ruleset, err := gameRuleset.LoadRuleset("./ruleset/default")
	if err != nil {
		log.Fatalf("failed to load ruleset: %v", err)
	}

	// Start the HTTP gameServer.
	address := fmt.Sprintf(":%s", PORT)

	server := gameServer.NewServer(address, &ruleset)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
