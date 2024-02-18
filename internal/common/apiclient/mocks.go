// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apiclient

import (
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/stretchr/testify/mock"
)

// MockAPIClient is a mock implementation of the IAPIClient interface, using testify's mock package.
type MockAPIClient struct {
	mock.Mock
}

// FetchUIConfigDTO mocks the FetchUIConfigDTO method of IAPIClient.
func (m *MockAPIClient) FetchUIConfigDTO() (*dtos.UIConfigDTO, error) {
	args := m.Called()
	return args.Get(0).(*dtos.UIConfigDTO), args.Error(1)
}
