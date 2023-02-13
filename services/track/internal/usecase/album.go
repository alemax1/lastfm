package usecase

import (
	"fmt"
	"track/internal/domain"
	"track/internal/repository"
)

type albumSearch struct {
	albumSearch repository.Album
}

func NewAlbumSearch(a repository.Album) AlbumSearch {
	return &albumSearch{
		albumSearch: a,
	}
}

type AlbumSearch interface {
	AlbumSearch(params map[string][]string) (domain.Response, error)
}

func (a albumSearch) AlbumSearch(params map[string][]string) (domain.Response, error) {
	album, err := a.albumSearch.AlbumSearch(params)
	if err != nil {
		return domain.Response{}, fmt.Errorf("albumSearch: %v", err)
	}

	return album, nil
}
