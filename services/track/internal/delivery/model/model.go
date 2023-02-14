package model

import "track/internal/domain"

type ErrorResponse struct {
	Error string `json:"error"`
}

type TracksResponse struct {
	Data []domain.TrackDBResponse `json:"data"`
}
