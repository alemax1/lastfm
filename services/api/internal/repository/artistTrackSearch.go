package repository

import (
	"api/internal/domain"
	"api/internal/repository/helper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

const (
	trackSearchURLTempl = "%s&api_key=%s"
)

type artistTrackSearch struct {
}

func NewArtistTrackSearch() ArtistTrackSearch {
	return new(artistTrackSearch)
}

type ArtistTrackSearch interface {
	GetArtistsAndTracks(params map[string][]string) ([]domain.Track, error)
}

func (a artistTrackSearch) GetArtistsAndTracks(params map[string][]string) ([]domain.Track, error) {
	trackSearchURL := fmt.Sprintf(trackSearchURLTempl, viper.GetString("lastfm.trackSearchURL"), viper.GetString("client.apiKey"))

	for k := range params {
		queryParams := url.Values{k: params[k]}

		trackSearchURL += "&" + queryParams.Encode()
	}

	resp, err := http.Get(trackSearchURL)
	if err != nil {
		return nil, fmt.Errorf("get: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %v", err)
	}

	var searchPage domain.ArtistTrackSearch

	if err := json.Unmarshal(respBytes, &searchPage); err != nil {
		return nil, fmt.Errorf("unmarshal: %v", err)
	}

	tracks, err := helper.EnrichTrackInfo(searchPage.Results.TrackMatches.Tracks)
	if err != nil {
		return nil, fmt.Errorf("enrich: %v", err)
	}

	return tracks, nil
}
