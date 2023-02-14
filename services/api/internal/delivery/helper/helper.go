package helper

import (
	"api/internal/delivery/model"
	"errors"
)

var RequiredError = errors.New("missing required parameter")

func CheckAlbumSearchParams(params model.AlbumSearchParams) error {
	if params.Album == "" {
		return RequiredError
	}

	if params.Artist == "" {
		return RequiredError
	}

	return nil
}
