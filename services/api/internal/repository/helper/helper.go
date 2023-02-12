package helper

import (
	"api/internal/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

const trackInfoURLTempl = "%s&api_key=%s&artist=%s&track=%s"

func EnrichTrackInfo(tracks []domain.Track) ([]domain.Track, error) {
	var trackInfo domain.TrackInfoSearch

	g := new(errgroup.Group)

	for i := range tracks {
		track := &tracks[i]

		g.Go(func() error {
			respBytes, err := getTrackInfo(track.Artist, track.Name)
			if err != nil {
				return fmt.Errorf("getTrackInfo: %v", err)
			}

			if err := json.Unmarshal(respBytes, &trackInfo); err != nil {
				return fmt.Errorf("unmarshal: %v", err)
			}

			track.Playcount = trackInfo.TrackInfo.Playcount
			track.Tags = trackInfo.TrackInfo.TopTags.Tags

			return nil
		})
	}
	g.Wait()

	return tracks, nil
}

func EnrichAlbumInfo(albumSearch domain.AlbumSearch) (domain.AlbumSearch, error) {
	var trackInfo domain.AlbumTrackSearch

	g := new(errgroup.Group)

	for i := range albumSearch.Album.Tracks {
		track := &albumSearch.Album.Tracks[i]

		g.Go(func() error {
			respBytes, err := getTrackInfo(albumSearch.Artist, albumSearch.Name)
			if err != nil {
				return fmt.Errorf("getTrackInfo: %v", err)
			}

			if err := json.Unmarshal(respBytes, &trackInfo); err != nil {
				return fmt.Errorf("unmarshal: %v", err)
			}

			track.Playcount = trackInfo.AlbumTrack.Playcount
			track.Tags = trackInfo.AlbumTrack.Tags
			track.Listeners = trackInfo.AlbumTrack.Listeners

			return nil
		})
	}
	g.Wait()

	return albumSearch, nil
}

func getTrackInfo(artist, name string) ([]byte, error) {
	trackInfoURL := fmt.Sprintf(trackInfoURLTempl,
		viper.GetString("lastfm.trackInfoURL"),
		viper.GetString("client.apiKey"),
		url.QueryEscape(artist),
		url.QueryEscape(name),
	)

	resp, err := http.Get(trackInfoURL)
	if err != nil {
		return nil, fmt.Errorf("get: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %v", err)
	}

	return respBytes, nil
}
