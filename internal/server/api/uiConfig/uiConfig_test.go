package uiConfig

import (
	"encoding/json"
	"github.com/hangovergames/eldoria/internal/common/dtos"
	"github.com/hangovergames/eldoria/internal/server/apiRequests"
	"github.com/hangovergames/eldoria/internal/server/apiResponses"
	"github.com/hangovergames/eldoria/internal/server/game/mocks"
	"testing"
)

func TestUIConfig(t *testing.T) {

	tests := []struct {
		name          string
		requestMethod bool // true for GET, false otherwise
		expectError   bool
	}{
		{"GET Request", true, false},
		{"Non-GET Request", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockResponse := &apiResponses.MockResponse{}
			mockRequest := &apiRequests.MockRequest{IsGet: tt.requestMethod}

			// Mock Ruleset
			mockRuleset := new(mocks.MockRuleset)
			uiConfigDTO := dtos.UIConfigDTO{
				SpriteSheets: []dtos.SpriteSheetDTO{
					{
						Name:        "Tiles",
						Image:       "freeciv/data/trident/tiles.png",
						TileWidth:   30,
						TileHeight:  30,
						TilesPerRow: 20,
						StartX:      0,
						StartY:      0,
					},
					{
						Name:        "OceanTiles",
						Image:       "freeciv/data/trident/tiles.png",
						TileWidth:   15,
						TileHeight:  15,
						TilesPerRow: 32,
						StartX:      0,
						StartY:      210,
					},
				},
				SpriteConfigs: []dtos.SpriteConfigDTO{
					{
						Name:      "ShallowOcean",
						SheetName: "OceanTiles",
						Index:     0,
					},
					{
						Name:      "DeepOcean",
						SheetName: "OceanTiles",
						Index:     10,
					},
					{
						Name:      "Grassland",
						SheetName: "Tiles",
						Index:     2,
					},
				},
				TileConfigs: []dtos.TileConfigDTO{
					{
						TileName: "DeepOcean",
						Sprites: []dtos.SpriteDTO{
							{Name: "DeepOcean", XOffset: 0, YOffset: 0},
							{Name: "DeepOcean", XOffset: 15, YOffset: 0},
							{Name: "DeepOcean", XOffset: 0, YOffset: 15},
							{Name: "DeepOcean", XOffset: 15, YOffset: 15},
						},
					},
					{
						TileName: "Grassland",
						Sprites: []dtos.SpriteDTO{
							{Name: "Grassland", XOffset: 0, YOffset: 0},
						},
					},
				},
			}
			mockRuleset.On("GetUI").Return(uiConfigDTO) // Setup expectation

			// Mock Server
			mockServer := new(mocks.MockServer)
			mockServer.On("GetRuleset").Return(mockRuleset) // Setup expectation

			UIConfig(mockResponse, mockRequest, mockServer)

			if tt.expectError {
				if !mockResponse.MethodNotAllowed {
					t.Errorf("Expected method not allowed error, but didn't get one")
				}
			} else {
				if mockResponse.SentStatusCode != 200 {
					t.Errorf("Expected status code 200, got %d", mockResponse.SentStatusCode)
				}
				// Assuming the SentData was marshaled into JSON, now unmarshal for assertion
				data, _ := json.Marshal(mockResponse.SentData) // Ignoring error for brevity
				var actualConfig dtos.UIConfigDTO
				if err := json.Unmarshal(data, &actualConfig); err != nil {
					t.Fatalf("Failed to unmarshal SentData: %v", err)
				}

				// Perform your assertions based on the actualConfig
				// Example assertion (adjust as needed)
				if len(actualConfig.SpriteSheets) == 0 || len(actualConfig.TileConfigs) == 0 {
					t.Errorf("Expected non-empty configurations for sprite sheets and tile configs")
				}

				// More detailed assertions can be added here based on the expected contents of your UIConfigDTO
			}
		})
	}
}
