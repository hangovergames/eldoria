// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameServer

import (
	"fmt"
	"github.com/hangovergames/eldoria/internal/server/api/index"
	"github.com/hangovergames/eldoria/internal/server/apiRequests"
	"github.com/hangovergames/eldoria/internal/server/apiResponses"
	"log"
	"net/http"
)

// Server represents the HTTP gameServer that handles the requests.
type Server struct {
	Address string
}

// HandlerFunc defines the type for handlers in this API.
type HandlerFunc func(apiResponses.Response, apiRequests.Request)

// NewServer creates and initializes a new Server instance.
func NewServer(address string) *Server {
	return &Server{
		Address: address,
	}
}

// SetupRoutes Define HTTP routes.
func (s *Server) SetupRoutes() {
	http.HandleFunc("/", responseHandler(index.Index))
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

// responseHandler wraps a handler function to inject dependencies.
func responseHandler(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := apiResponses.NewJSONResponse(w)
		request := apiRequests.NewRequest(r)
		handler(response, request)
	}
}
