package model

import "api/internal/domain"

type TrackAndArtistsResponse struct {
	Data []domain.Track `json:"data"`
}

type AlbumSearchResponse struct {
	Data domain.Album `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
