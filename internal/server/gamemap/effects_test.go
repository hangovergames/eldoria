// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gamemap

import (
	"github.com/hangovergames/eldoria/internal/server/game"
	"testing"
)

func TestTileEffectToString(t *testing.T) {
	tests := []struct {
		effect game.TileEffect
		want   string
	}{
		{game.Passable, "Passable"},
		{game.ProvidesFreshWater, "ProvidesFreshWater"},
		{0, "Unknown"}, // Testing an undefined effect
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := TileEffectToString(tt.effect)
			if got != tt.want {
				t.Errorf("TileEffectToString(%v) = %v, want %v", tt.effect, got, tt.want)
			}
		})
	}
}

func TestStringToTileEffect(t *testing.T) {
	tests := []struct {
		effectStr string
		want      game.TileEffect
		wantErr   bool
	}{
		{"Passable", game.Passable, false},
		{"ProvidesFreshWater", game.ProvidesFreshWater, false},
		{"NonexistentEffect", 0, true}, // Testing an unknown effect
	}

	for _, tt := range tests {
		t.Run(tt.effectStr, func(t *testing.T) {
			got, err := StringToTileEffect(tt.effectStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToTileEffect(%v) error = %v, wantErr %v", tt.effectStr, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToTileEffect(%v) = %v, want %v", tt.effectStr, got, tt.want)
			}
		})
	}
}

func TestCombineTileEffects(t *testing.T) {
	tests := []struct {
		name  string
		input []game.TileEffect
		want  game.TileEffect
	}{
		{"PassableAndFertile", []game.TileEffect{game.Passable, game.Fertile}, game.Passable | game.Fertile},
		{"Empty", []game.TileEffect{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombineTileEffects(tt.input); got != tt.want {
				t.Errorf("CombineTileEffects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringsToTileEffects(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		wantErr   bool
		wantCount int // Using count to simplify checking
	}{
		{"ValidEffects", []string{"Passable", "Fertile"}, false, 2},
		{"IncludesUnknown", []string{"Passable", "NonexistentEffect"}, true, 0},
		{"Empty", []string{}, false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringsToTileEffects(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringsToTileEffects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantCount {
				t.Errorf("StringsToTileEffects() got %v effects, want %v", len(got), tt.wantCount)
			}
		})
	}
}

func TestTileEffect_HasTileEffect(t *testing.T) {
	tests := []struct {
		name       string
		tileEffect game.TileEffect
		effect     game.TileEffect
		want       bool
	}{
		{
			name:       "SingleEffectTrue",
			tileEffect: game.Passable,
			effect:     game.Passable,
			want:       true,
		},
		{
			name:       "SingleEffectFalse",
			tileEffect: game.Passable,
			effect:     game.Fertile,
			want:       false,
		},
		{
			name:       "CombinedEffectTrue",
			tileEffect: game.Passable | game.Fertile,
			effect:     game.Fertile,
			want:       true,
		},
		{
			name:       "CombinedEffectFalse",
			tileEffect: game.Passable | game.Fertile,
			effect:     game.RichInWood,
			want:       false,
		},
		{
			name:       "NoEffect",
			tileEffect: 0,
			effect:     game.Passable,
			want:       false,
		},
		{
			name:       "CheckAgainstNoEffect",
			tileEffect: game.Passable,
			effect:     0,
			want:       false, // Depending on interpretation, checking against "no effect" could be true or false. Adjust based on your logic.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tileEffect.HasTileEffect(tt.effect); got != tt.want {
				t.Errorf("TileEffect.HasTileEffect(%v) = %v, want %v", tt.effect, got, tt.want)
			}
		})
	}
}
