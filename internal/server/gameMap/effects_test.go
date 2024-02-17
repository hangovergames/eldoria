// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package gameMap

import "testing"

func TestTileEffectToString(t *testing.T) {
	tests := []struct {
		effect TileEffect
		want   string
	}{
		{Passable, "Passable"},
		{ProvidesFreshWater, "ProvidesFreshWater"},
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
		want      TileEffect
		wantErr   bool
	}{
		{"Passable", Passable, false},
		{"ProvidesFreshWater", ProvidesFreshWater, false},
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
		input []TileEffect
		want  TileEffect
	}{
		{"PassableAndFertile", []TileEffect{Passable, Fertile}, Passable | Fertile},
		{"Empty", []TileEffect{}, 0},
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
		tileEffect TileEffect
		effect     TileEffect
		want       bool
	}{
		{
			name:       "SingleEffectTrue",
			tileEffect: Passable,
			effect:     Passable,
			want:       true,
		},
		{
			name:       "SingleEffectFalse",
			tileEffect: Passable,
			effect:     Fertile,
			want:       false,
		},
		{
			name:       "CombinedEffectTrue",
			tileEffect: Passable | Fertile,
			effect:     Fertile,
			want:       true,
		},
		{
			name:       "CombinedEffectFalse",
			tileEffect: Passable | Fertile,
			effect:     RichInWood,
			want:       false,
		},
		{
			name:       "NoEffect",
			tileEffect: 0,
			effect:     Passable,
			want:       false,
		},
		{
			name:       "CheckAgainstNoEffect",
			tileEffect: Passable,
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
