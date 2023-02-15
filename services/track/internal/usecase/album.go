package usecase

import (
	"context"
	"track/internal/domain"
	"track/internal/repository"
)

type albumSearch struct {
	albumPG repository.AlbumPG
}

func NewAlbumSearch(aPG repository.AlbumPG) AlbumSearch {
	return &albumSearch{
		albumPG: aPG,
	}
}

type AlbumSearch interface {
	AlbumSearch(ctx context.Context, params map[string][]string, album domain.AlbumResponse) error
}

func (a albumSearch) AlbumSearch(ctx context.Context, params map[string][]string, album domain.AlbumResponse) error {
	if err := a.albumPG.SaveAlbumInfoToDB(ctx, album); err != nil {
		return err
	}

	return nil
}
