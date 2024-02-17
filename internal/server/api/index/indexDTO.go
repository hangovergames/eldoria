// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package index

type IndexDTO struct {
	Version string `json:"version"`
}

func newIndexDTO(version string) IndexDTO {
	return IndexDTO{
		Version: version,
	}
}
