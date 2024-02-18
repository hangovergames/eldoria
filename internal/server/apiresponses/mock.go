// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apiresponses

// MockResponse implements the apiResponses.Response interface
type MockResponse struct {
	SentData         interface{}
	SentStatusCode   int
	SentErrorMessage string
	MethodNotAllowed bool
}

func (m *MockResponse) Send(statusCode int, data interface{}) {
	m.SentStatusCode = statusCode
	m.SentData = data
}

func (m *MockResponse) SendError(statusCode int, errorMessage string) {
	m.SentStatusCode = statusCode
	m.SentErrorMessage = errorMessage
}

func (m *MockResponse) SendMethodNotSupportedError() {
	m.MethodNotAllowed = true
}
