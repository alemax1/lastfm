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

type artistTrackSearchRepository struct {
}

func NewArtistTrackSearchRepository() domain.ArtistTrackSearchRepository {
	return new(artistTrackSearchRepository)
}

func (a artistTrackSearchRepository) GetArtistsAndTracksByPageAndLimit(params map[string][]string) ([]domain.Track, error) {
	trackSearchURL := fmt.Sprintf("%s&api_key=%s", viper.GetString("lastfm.trackSearchURL"), viper.GetString("client.apiKey"))

	for k := range params {
		for i := range params[k] {
			paramWithoutSpaces := url.QueryEscape(params[k][i])

			trackSearchURL += fmt.Sprintf("&%s=%s", k, paramWithoutSpaces)
		}
	}

	resp, err := http.Get(trackSearchURL)
	if err != nil {
		return nil, errors.Wrap(err, "get track")
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}
	defer resp.Body.Close()

	var searchPage domain.ArtistTrackSearch

	if err := json.Unmarshal(respBytes, &searchPage); err != nil {
		return nil, errors.Wrap(err, "unmarshal page")
	}

	tracks, err := helper.TrackInfoSearch(searchPage.Results.TrackMatches.Tracks)
	if err != nil {
		return nil, errors.Wrap(err, "helper")
	}

	return tracks, nil
}
