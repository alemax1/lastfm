package repository

import (
	"api/internal/domain"
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

	for i := range searchPage.Results.TrackMatches.Tracks {
		trackInfoURL := fmt.Sprintf("%s&api_key=%s&artist=%s&track=%s",
			viper.GetString("lastfm.trackInfoURL"),
			viper.GetString("client.apiKey"),
			url.QueryEscape(searchPage.Results.TrackMatches.Tracks[i].Artist),
			url.QueryEscape(searchPage.Results.TrackMatches.Tracks[i].Name),
		)

		resp, err := http.Get(trackInfoURL)
		if err != nil {
			return nil, errors.Wrap(err, "get info")
		}

		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "read body")
		}
		defer resp.Body.Close()

		var trackInfo domain.TrackInfoSearch

		if err := json.Unmarshal(respBytes, &trackInfo); err != nil {
			return nil, errors.Wrap(err, "unmarshal info")
		}

		searchPage.Results.TrackMatches.Tracks[i].Playcount = trackInfo.TrackInfo.Playcount
		searchPage.Results.TrackMatches.Tracks[i].Tags = trackInfo.TrackInfo.TopTags.Tag
	}

	return searchPage.Results.TrackMatches.Tracks, nil
}
