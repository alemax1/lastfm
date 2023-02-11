package usecase

import (
	"api/internal/domain"
)

type albumSearchUsecase struct {
	albumSearchRepository domain.AlbumSearchRepository
}

func NewAlbumSearchUsecase(a domain.AlbumSearchRepository) domain.AlbumSearchUsecase {
	return &albumSearchUsecase{
		albumSearchRepository: a,
	}
}

func (a albumSearchUsecase) GetAlbumInfoByTitleAndArtist(params map[string][]string) (*domain.Album, error) {
	album, err := a.albumSearchRepository.GetAlbumInfoByTitleAndArtist(params)
	if err != nil {
		return nil, err
	}

	return album, nil
}
