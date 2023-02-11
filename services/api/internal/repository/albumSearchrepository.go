package repository

import (
	"api/internal/domain"
	"api/internal/repository/helper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type albumSearchRepository struct {
}

func NewAlbumSearchRepository() domain.AlbumSearchRepository {
	return new(albumSearchRepository)
}

func (a albumSearchRepository) GetAlbumInfoByTitleAndArtist(params map[string][]string) (*domain.Album, error) {
	albumSearchURL := fmt.Sprintf("%s&api_key=%s", viper.GetString("lastfm.albumSearchURL"), viper.GetString("client.apiKey"))

	for k := range params {
		for i := range params[k] {
			paramWithoutSpaces := url.QueryEscape(params[k][i])

			albumSearchURL += fmt.Sprintf("&%s=%s", k, paramWithoutSpaces)
		}
	}

	resp, err := http.Get(albumSearchURL)
	if err != nil {
		return nil, errors.Wrap(err, "get")
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}
	defer resp.Body.Close()

	fmt.Println(string(respBytes))

	var album domain.AlbumSearch

	if err := json.Unmarshal(respBytes, &album); err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	albumTracks, err := helper.AlbumTrackInfoSearch(album)

	return &albumTracks.Album, nil
}
