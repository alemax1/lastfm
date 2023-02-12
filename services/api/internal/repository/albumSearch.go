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
	albumSearchURLTempl = "%s&api_key=%s"
)

type albumSearch struct {
}

func NewAlbumSearch() AlbumSearch {
	return new(albumSearch)
}

type AlbumSearch interface {
	GetAlbumInfoByTitleAndArtist(params map[string][]string) (domain.Album, error)
}

func (a albumSearch) GetAlbumInfoByTitleAndArtist(params map[string][]string) (domain.Album, error) {
	albumSearchURL := fmt.Sprintf(albumSearchURLTempl, viper.GetString("lastfm.albumSearchURL"), viper.GetString("client.apiKey"))

	for k := range params {
		queryParams := url.Values{k: params[k]}

		albumSearchURL += "&" + queryParams.Encode()
	}

	resp, err := http.Get(albumSearchURL)
	if err != nil {
		return domain.Album{}, fmt.Errorf("get: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return domain.Album{}, fmt.Errorf("read body: %v", err)
	}

	var album domain.AlbumSearch

	if err := json.Unmarshal(respBytes, &album); err != nil {
		return domain.Album{}, fmt.Errorf("unmarshal: %v", err)
	}

	albumTracks, err := helper.EnrichAlbumInfo(album)
	if err != nil {
		return domain.Album{}, fmt.Errorf("enrich: %v", err)
	}

	return albumTracks.Album, nil
}
