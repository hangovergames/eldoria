// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameServer

import (
	"fmt"
	"github.com/hangovergames/eldoria/internal/server/api/index"
	"github.com/hangovergames/eldoria/internal/server/api/uiConfig"
	"github.com/hangovergames/eldoria/internal/server/apiRequests"
	"github.com/hangovergames/eldoria/internal/server/apiResponses"
	"github.com/hangovergames/eldoria/internal/server/game"
	"log"
	"net/http"
)

// Server represents the HTTP gameServer that handles the requests.
type Server struct {
	Address string
	Ruleset game.IRuleset
}

// HandlerFunc defines the type for handlers in this API.
type HandlerFunc func(apiResponses.Response, apiRequests.Request, game.IServer)

// NewServer creates and initializes a new Server instance.
func NewServer(
	address string,
	ruleset game.IRuleset,
) *Server {
	return &Server{
		Address: address,
		Ruleset: ruleset,
	}
}

// SetupRoutes Define HTTP routes.
func (s *Server) SetupRoutes() {
	http.HandleFunc("/", responseHandler(index.Index, s))
	http.HandleFunc("/ui/config", responseHandler(uiConfig.UIConfig, s))
}

// Start begins listening on the specified port and starts handling incoming requests.
func (s *Server) Start() error {
	s.SetupRoutes()
	log.Printf("Starting gameServer at %s", s.Address)
	if err := http.ListenAndServe(s.Address, nil); err != nil {
		return fmt.Errorf("could not start gameServer for %s: %w", s.Address, err)
	}
	return nil
}

func (s *Server) GetAddress() string {
	return s.Address
}

func (s *Server) GetRuleset() game.IRuleset {
	return s.Ruleset
}

// responseHandler wraps a handler function to inject dependencies.
func responseHandler(
	handler HandlerFunc,
	server game.IServer,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := apiResponses.NewJSONResponse(w)
		request := apiRequests.NewRequest(r)
		handler(response, request, server)
	}
}
