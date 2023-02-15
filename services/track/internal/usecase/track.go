package usecase

import (
	"context"
	"track/internal/domain"
	"track/internal/repository"
)

type track struct {
	trackPG repository.TrackPG
}

func NewTrack(tPG repository.TrackPG) Track {
	return &track{
		trackPG: tPG,
	}
}

type Track interface {
	TrackSearch(ctx context.Context, params map[string][]string, tracks domain.TracksResponse) error
	GetTracksByTag(ctx context.Context, page, limit int, tag string) ([]domain.TrackDBResponse, error)
	GetTracksByArtist(ctx context.Context, page, limit int, artist string) ([]domain.TrackDBResponse, error)
}

func (t track) TrackSearch(ctx context.Context, params map[string][]string, tracks domain.TracksResponse) error {
	if err := t.trackPG.SaveTrackInfoToDB(ctx, tracks); err != nil {
		return err
	}

	return nil
}

func (t track) GetTracksByTag(ctx context.Context, page, limit int, tag string) ([]domain.TrackDBResponse, error) {
	tracks, err := t.trackPG.GetTracksByTag(ctx, page, limit, tag)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (t track) GetTracksByArtist(ctx context.Context, page, limit int, artist string) ([]domain.TrackDBResponse, error) {
	tracks, err := t.trackPG.GetTracksByArtist(ctx, page, limit, artist)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}
