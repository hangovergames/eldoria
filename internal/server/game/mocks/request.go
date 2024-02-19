// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import "net/url"

// MockRequest implements the apiRequests.IRequest interface
type MockRequest struct {
	IsGet bool
	URL   *url.URL
	Vars  map[string]string
}

func (m *MockRequest) IsMethodGet() bool {
	return m.IsGet
}

func (m *MockRequest) GetURL() *url.URL {
	return m.URL
}

func (m *MockRequest) GetVars() map[string]string {
	return m.Vars
}
