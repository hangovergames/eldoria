// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameserver

import (
	"github.com/hangovergames/eldoria/internal/server/api/index"
	"github.com/hangovergames/eldoria/internal/server/game/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {

	ruleset := mocks.NewMockRuleset()

	mockGameState := new(mocks.MockGameState)
	mockMap := new(mocks.MockGameMap)
	mockMap.On("GetWidth").Return(10)
	//mockMap.On("GetHeight").Return(10)
	//mockPlayers := []gamePlayer.Player{{ /* setup player data */ }}

	// Setup expectations
	//mockGameState.On("GetMap").Return(mockMap)
	//mockGameState.On("GetPlayers").Return(mockPlayers)

	s := NewServer(":8080", ruleset, mockGameState)
	go func() {
		if err := s.Start(); err != nil {
			t.Errorf("Failed to start gameServer: %v", err)
		}
	}()

	// Wait a moment for the gameServer to start
	time.Sleep(time.Second)

	// Make a request to verify the gameServer is responding
	res, err := http.Get("http://" + s.Address)
	if err != nil {
		t.Fatalf("Failed to make request to gameServer: %v", err)
	}
	defer res.Body.Close()

	// Check for a successful response
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}

	// Additional checks can be added here to verify the gameServer's behavior

	mockGameState.AssertExpectations(t)

}

func TestResponseHandler(t *testing.T) {

	// Create a MockServer.
	mockServer := &mocks.MockServer{}

	// If you're testing handlers that use the ruleset, provide a mock IRuleset too.
	mockRuleset := &mocks.MockRuleset{}
	mockServer.Ruleset = mockRuleset

	// Handler that we want to test.
	handler := responseHandler(index.Index, mockServer)

	// Create a test server with the handler.
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// Make a request to the test server.
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer res.Body.Close()

	// Verify that your server methods were called as expected.
	mockServer.AssertExpectations(t)
	// Here you might also want to verify that the ruleset methods were called, if applicable.
	mockRuleset.AssertExpectations(t)

	// Further assertions can be made based on the response.
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got: %v", res.Status)
	}
	// Additional assertions can be made on the response body if necessary.

}
