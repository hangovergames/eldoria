// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package dtos

type PlayerDTO struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	Score int    `json:"score"`
}

func NewPlayerDTO(
	name string,
	level, score int,
) PlayerDTO {
	return PlayerDTO{
		Name:  name,
		Level: level,
		Score: score,
	}
}
