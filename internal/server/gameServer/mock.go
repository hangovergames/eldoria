// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameServer

import (
	"github.com/hangovergames/eldoria/internal/server/apiRequests"
	"github.com/hangovergames/eldoria/internal/server/apiResponses"
)

// mockHandler is a test utility that implements the HandlerFunc signature.
// It allows us to inspect the arguments passed to it during the test.
type mockHandler struct {
	called   bool
	request  apiRequests.Request
	response apiResponses.Response
}

func (m *mockHandler) Serve(response apiResponses.Response, request apiRequests.Request) {
	m.called = true
	m.request = request
	m.response = response
}
