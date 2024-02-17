// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameServer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {

	s := NewServer(":8080")
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

}

func TestResponseHandler(t *testing.T) {
	mh := &mockHandler{}
	handler := responseHandler(mh.Serve)

	// Create a test gameServer with the handler.
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// Make a request to the test gameServer.
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer res.Body.Close()

	// Verify the handler was called.
	if !mh.called {
		t.Errorf("Expected mock handler to be called")
	}

	// Here you might also want to verify the response and request objects
	// However, since they are interfaces, you'd typically check for expected behavior
	// rather than direct object inspection. This might involve mocking the Response
	// and Request interfaces to track method calls or inspect passed values.
}
