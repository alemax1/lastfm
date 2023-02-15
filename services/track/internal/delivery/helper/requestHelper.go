package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"track/internal/domain"

	"github.com/spf13/viper"
)

func AlbumSearch(params url.Values) (domain.AlbumResponse, error) {
	albumSearchURL := fmt.Sprintf(
		"%s:%d/%s?",
		viper.GetString("trackAggregator.host"),
		viper.GetInt("trackAggregator.port"),
		viper.GetString("trackAggregator.methods.album"),
	)

	for k := range params {
		queryParams := url.Values{k: params[k]}

		albumSearchURL += "&" + queryParams.Encode()
	}

	resp, err := http.Get(albumSearchURL)
	if err != nil {
		return domain.AlbumResponse{}, fmt.Errorf("get: %v", err)
	}
	defer resp.Body.Close()

	var album domain.AlbumResponse

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return domain.AlbumResponse{}, fmt.Errorf("read body: %v", err)
	}

	if err := json.Unmarshal(respBytes, &album); err != nil {
		return domain.AlbumResponse{}, fmt.Errorf("unmarshal: %v", err)
	}

	return album, nil
}

func TrackSearch(params map[string][]string) (domain.TracksResponse, error) {
	trackSearchURL := fmt.Sprintf(
		"%s:%d/%s?",
		viper.GetString("trackAggregator.host"),
		viper.GetInt("trackAggregator.port"),
		viper.GetString("trackAggregator.methods.tracks"),
	)

	for k := range params {
		queryParams := url.Values{k: params[k]}

		trackSearchURL += "&" + queryParams.Encode()
	}

	resp, err := http.Get(trackSearchURL)
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
