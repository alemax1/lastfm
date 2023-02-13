package usecase

import (
	"fmt"
	"track/internal/domain"
	"track/internal/repository"
)

type trackSearch struct {
	trackSearch repository.Track
	trackPG     repository.TrackPG
}

func NewTrack(t repository.Track, tPG repository.TrackPG) TrackSearch {
	return &trackSearch{
		trackSearch: t,
		trackPG:     tPG,
	}
}

type TrackSearch interface {
	TrackSearch(params map[string][]string) (domain.TracksResponse, error)
}

func (t trackSearch) TrackSearch(params map[string][]string) (domain.TracksResponse, error) {
	tracks, err := t.trackSearch.TrackSearch(params)
	if err != nil {
		return domain.TracksResponse{}, fmt.Errorf("trackSearch: %v", err)
	}

	if err := t.trackPG.SaveTrackInfoToDB(tracks); err != nil {
		return domain.TracksResponse{}, fmt.Errorf("save: %v", err)
	}

	return tracks, nil
}
