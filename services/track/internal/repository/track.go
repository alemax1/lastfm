package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"track/internal/domain"

	"github.com/spf13/viper"
)

type track struct {
}

func NewTrack() Track {
	return new(track)
}

type Track interface {
	TrackSearch(params map[string][]string) (domain.TracksResponse, error)
}

func (t track) TrackSearch(params map[string][]string) (domain.TracksResponse, error) {
	trackSearchURL := fmt.Sprintf(
		"%s:%d/%s",
		viper.GetString("trackAggregator.host"),
		viper.GetInt("trackAggregator.port"),
		viper.GetString("trackAggregator.methods.tracks"),
	)

	URLWithQuery, err := url.Parse(trackSearchURL)
	if err != nil {
		return domain.TracksResponse{}, fmt.Errorf("parse: %v", err)
	}

	values := URLWithQuery.Query()

	for k := range params {
		for i := range params[k] {
			values.Add(k, params[k][i])
		}
	}

	URLWithQuery.RawQuery = values.Encode()

	resp, err := http.Get(URLWithQuery.String())
	if err != nil {
		return domain.TracksResponse{}, fmt.Errorf("get: %v", err)
	}
	defer resp.Body.Close()

	var tracks domain.TracksResponse

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return domain.TracksResponse{}, fmt.Errorf("read body: %v", err)
	}

	if err := json.Unmarshal(respBytes, &tracks); err != nil {
		return domain.TracksResponse{}, fmt.Errorf("unmarshal: %v", err)
	}

	return tracks, nil
}
