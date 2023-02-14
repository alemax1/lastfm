package usecase

import (
	"fmt"
	"track/internal/domain"
	"track/internal/repository"
)

type albumSearch struct {
	albumSearch repository.Album
	albumPG     repository.AlbumPG
}

func NewAlbumSearch(a repository.Album, aPG repository.AlbumPG) AlbumSearch {
	return &albumSearch{
		albumSearch: a,
		albumPG:     aPG,
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

	if err := a.albumPG.SaveAlbumInfoToDB(album.Data); err != nil {
		return domain.Response{}, fmt.Errorf("save: %v", err)
	}

	return album, nil
}
