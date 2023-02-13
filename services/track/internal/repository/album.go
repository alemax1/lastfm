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

type album struct{}

func NewAlbum() Album {
	return new(album)
}

type Album interface {
	AlbumSearch(params map[string][]string) (domain.Response, error)
}

func (a album) AlbumSearch(params map[string][]string) (domain.Response, error) {
	albumSearchURL := fmt.Sprintf(
		"%s:%d/%s",
		viper.GetString("trackAggregator.host"),
		viper.GetInt("trackAggregator.port"),
		viper.GetString("trackAggregator.methods.album"),
	)

	URLWithQuery, err := url.Parse(albumSearchURL)
	if err != nil {
		return domain.Response{}, fmt.Errorf("parse: %v", err)
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
		return domain.Response{}, fmt.Errorf("get: %v", err)
	}
	defer resp.Body.Close()

	var album domain.Response

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return domain.Response{}, fmt.Errorf("read body: %v", err)
	}

	if err := json.Unmarshal(respBytes, &album); err != nil {
		return domain.Response{}, fmt.Errorf("unmarshal: %v", err)
	}

	return album, nil
}
