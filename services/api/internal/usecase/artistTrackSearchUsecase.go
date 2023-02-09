package usecase

import "api/internal/domain"

type artistTrackSearchUsecase struct {
	artistTrackSearchRepository domain.ArtistTrackSearchRepository
}

func NewArtistTrackSearchUsecase(a domain.ArtistTrackSearchRepository) domain.ArtistTrackSearchUsecase {
	return &artistTrackSearchUsecase{
		artistTrackSearchRepository: a,
	}
}

func (a artistTrackSearchUsecase) GetAllArtistsAndTracks() ([]domain.Track, error) {
	tracks, err := a.artistTrackSearchRepository.GetAllArtistsAndTracks()
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (a artistTrackSearchUsecase) GetArtistsAndTracksByPageAndLimit(limit, page int) ([]domain.Track, error) {
	return nil, nil
}
