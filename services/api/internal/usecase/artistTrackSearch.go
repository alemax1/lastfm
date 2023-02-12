package usecase

import (
	"api/internal/domain"
	"api/internal/repository"
)

type artistTrackSearch struct {
	artistTrackSearch ArtistTrackSearch
}

func NewArtistTrackSearch(a repository.ArtistTrackSearch) ArtistTrackSearch {
	return &artistTrackSearch{
		artistTrackSearch: a,
	}
}

type ArtistTrackSearch interface {
	GetArtistsAndTracks(params map[string][]string) ([]domain.Track, error)
}

func (a artistTrackSearch) GetArtistsAndTracks(params map[string][]string) (tracks []domain.Track, err error) {
	tracks, err = a.artistTrackSearch.GetArtistsAndTracks(params)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}
