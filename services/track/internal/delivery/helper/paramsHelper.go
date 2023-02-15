package helper

import (
	"net/url"
	"track/internal/domain"
)

func CheckTrackSearchParams(params url.Values) error {
	track, ok := params["track"]
	if !ok {
		return domain.RequiredError
	}

	artist, ok := params["artist"]
	if !ok {
		return domain.RequiredError
	}

	if track[0] == "" || artist[0] == "" {
		return domain.RequiredError
	}

	return nil
}

func CheckTracksByTagParams(params url.Values) (string, error) {
	if _, ok := params["tag"]; !ok {
		return "", domain.RequiredError
	}

	tag := params.Get("tag")

	if tag == "" {
		return "", domain.RequiredError
	}

	return tag, nil
}

func CheckTracksByArtistParams(params url.Values) (string, error) {
	if _, ok := params["artist"]; !ok {
		return "", domain.RequiredError
	}

	artist := params.Get("artist")

	if artist == "" {
		return "", domain.RequiredError
	}

	return artist, nil
}

func CheckAlbumSearchParams(params url.Values) error {
	album, ok := params["album"]
	if !ok {
		return domain.RequiredError
	}

	artist, ok := params["artist"]
	if !ok {
		return domain.RequiredError
	}

	if album[0] == "" || artist[0] == "" {
		return domain.RequiredError
	}

	return nil
}
