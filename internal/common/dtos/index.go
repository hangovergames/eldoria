// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package dtos

type IndexDTO struct {
	Version string `json:"version"`
}

func NewIndexDTO(version string) IndexDTO {
	return IndexDTO{
		Version: version,
	}
}
