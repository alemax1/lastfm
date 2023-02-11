package helper

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

func TrackInfoSearch(tracks []domain.Track) ([]domain.Track, error) {
	for i := range tracks {
		trackInfoURL := fmt.Sprintf("%s&api_key=%s&artist=%s&track=%s",
			viper.GetString("lastfm.trackInfoURL"),
			viper.GetString("client.apiKey"),
			url.QueryEscape(tracks[i].Artist),
			url.QueryEscape(tracks[i].Name),
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

		tracks[i].Playcount = trackInfo.TrackInfo.Playcount
		tracks[i].Tags = trackInfo.TrackInfo.TopTags.Tag
	}

	return tracks, nil
}

func AlbumTrackInfoSearch(albumSearch domain.AlbumSearch) (*domain.AlbumSearch, error) {
	for i := range albumSearch.Album.Tracks {
		trackInfoURL := fmt.Sprintf("%s&api_key=%s&artist=%s&track=%s",
			viper.GetString("lastfm.trackInfoURL"),
			viper.GetString("client.apiKey"),
			url.QueryEscape(albumSearch.Album.Artist),
			url.QueryEscape(albumSearch.Album.Tracks[i].Name),
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

		albumSearch.Album.Tracks[i].Playcount = trackInfo.Track.Playcount
		albumSearch.Album.Tracks[i].Tags = trackInfo.Track.Tags
		albumSearch.Album.Tracks[i].Listeners = trackInfo.Track.Listeners
	}

	return &albumSearch, nil
}
