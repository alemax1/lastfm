package usecase

import (
	"api/internal/domain"
)

type artistTrackSearchUsecase struct {
	artistTrackSearchRepository domain.ArtistTrackSearchRepository
}

func NewArtistTrackSearchUsecase(a domain.ArtistTrackSearchRepository) domain.ArtistTrackSearchUsecase {
	return &artistTrackSearchUsecase{
		artistTrackSearchRepository: a,
	}
}

func (a artistTrackSearchUsecase) GetArtistsAndTracksByPageAndLimit(params map[string][]string) (tracks []domain.Track, err error) {
	tracks, err = a.artistTrackSearchRepository.GetArtistsAndTracksByPageAndLimit(params)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}
