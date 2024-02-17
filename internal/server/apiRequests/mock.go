// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apiRequests

// MockRequest implements the apiRequests.Request interface
type MockRequest struct {
	IsGet bool
}

func (m *MockRequest) IsMethodGet() bool {
	return m.IsGet
}
