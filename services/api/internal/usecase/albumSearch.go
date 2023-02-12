package usecase

import (
	"api/internal/domain"
	"api/internal/repository"
)

type albumSearch struct {
	albumSearch repository.AlbumSearch
}

func NewAlbumSearch(a repository.AlbumSearch) AlbumSearch {
	return &albumSearch{
		albumSearch: a,
	}
}

type AlbumSearch interface {
	GetAlbumByTitleAndArtist(params map[string][]string) (domain.Album, error)
}

func (a albumSearch) GetAlbumByTitleAndArtist(params map[string][]string) (domain.Album, error) {
	album, err := a.albumSearch.GetAlbumInfoByTitleAndArtist(params)
	if err != nil {
		return domain.Album{}, err
	}

	return album, nil
}
