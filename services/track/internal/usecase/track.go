package usecase

import (
	"fmt"
	"track/internal/domain"
	"track/internal/repository"
)

type track struct {
	trackSearch repository.Track
	trackPG     repository.TrackPG
}

func NewTrack(t repository.Track, tPG repository.TrackPG) Track {
	return &track{
		trackSearch: t,
		trackPG:     tPG,
	}
}

type Track interface {
	TrackSearch(params map[string][]string) (domain.TracksResponse, error)
	GetTracksByTag(page, limit int, tag string) ([]domain.TrackDBResponse, error)
	GetTracksByArtist(page, limit int, artist string) ([]domain.TrackDBResponse, error)
}

func (t track) TrackSearch(params map[string][]string) (domain.TracksResponse, error) {
	tracks, err := t.trackSearch.TrackSearch(params)
	if err != nil {
		return domain.TracksResponse{}, fmt.Errorf("trackSearch: %v", err)
	}

	if err := t.trackPG.SaveTrackInfoToDB(tracks); err != nil {
		return domain.TracksResponse{}, fmt.Errorf("save: %v", err)
	}

	return tracks, nil
}

func (t track) GetTracksByTag(page, limit int, tag string) ([]domain.TrackDBResponse, error) {
	tracks, err := t.trackPG.GetTracksByTag(page, limit, tag)
	if err != nil {
		return nil, fmt.Errorf("getTracks: %v", err)
	}

	return tracks, nil
}

func (t track) GetTracksByArtist(page, limit int, artist string) ([]domain.TrackDBResponse, error) {
	tracks, err := t.trackPG.GetTracksByArtist(page, limit, artist)
	if err != nil {
		return nil, fmt.Errorf("tracks artist: %v", err)
	}

	return tracks, nil
}
