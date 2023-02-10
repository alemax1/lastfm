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

	var album domain.AlbumSearch

	if err := json.Unmarshal(respBytes, &album); err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	for i := range album.Album.Tracks.Tracks {
		trackInfoURL := fmt.Sprintf("%s&api_key=%s&artist=%s&track=%s",
			viper.GetString("lastfm.trackInfoURL"),
			viper.GetString("client.apiKey"),
			url.QueryEscape(album.Album.Artist),
			url.QueryEscape(album.Album.Tracks.Tracks[i].Name),
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

		var trackInfo domain.AlbumTrackSearch

		if err := json.Unmarshal(respBytes, &trackInfo); err != nil {
			return nil, errors.Wrap(err, "unmarshal info")
		}

		album.Album.Tracks.Tracks[i].Playcount = trackInfo.Track.Playcount
		album.Album.Tracks.Tracks[i].Tags = trackInfo.Track.Tags
		album.Album.Tracks.Tracks[i].Listeners = trackInfo.Track.Listeners
	}

	return &album.Album, nil
}
