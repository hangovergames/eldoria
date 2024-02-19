// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hangovergames/eldoria/internal/server/api/index"
	"github.com/hangovergames/eldoria/internal/server/api/players"
	"github.com/hangovergames/eldoria/internal/server/api/uiconfig"
	"github.com/hangovergames/eldoria/internal/server/apirequests"
	"github.com/hangovergames/eldoria/internal/server/apiresponses"
	"github.com/hangovergames/eldoria/internal/server/game"
	"log"
	"net/http"
)

// Server represents the HTTP gameServer that handles the requests.
type Server struct {
	Address string
	Ruleset game.IRuleset
	State   game.IGameState
}

// NewServer creates and initializes a new Server instance.
func NewServer(
	address string,
	ruleset game.IRuleset,
	state game.IGameState,
) *Server {
	return &Server{
		Address: address,
		Ruleset: ruleset,
		State:   state,
	}
}

// SetupRoutes Define HTTP routes.
func (s *Server) SetupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/players/{username:[a-z0-9_]+}", responseHandler(players.PlayerCheck, s)).Methods("GET")
	r.HandleFunc("/ui/config", responseHandler(uiconfig.UIConfig, s)).Methods("GET")
	r.HandleFunc("/", responseHandler(index.Index, s)).Methods("GET")
	http.Handle("/", r)
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

func (s *Server) GetState() game.IGameState {
	return s.State
}

func (s *Server) GetRuleset() game.IRuleset {
	return s.Ruleset
}

// responseHandler wraps a handler function to inject dependencies.
func responseHandler(
	handler game.RequestHandlerFunc,
	server game.IServer,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := apiresponses.NewJSONResponse(w)
		request := apirequests.NewRequest(r)
		handler(response, request, server)
	}
}
