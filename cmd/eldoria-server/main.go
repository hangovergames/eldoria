// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package main

import (
	"fmt"
	server2 "github.com/hangovergames/eldoria/internal/server"
	"log"
)

const PORT = "8080"

func main() {

	// Start the HTTP server.
	address := fmt.Sprintf(":%s", PORT)

	server := server2.NewServer(address)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
