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

const (
	albumSearchURLTempl = "%s&api_key=%s"
	trackInfoURLTempl   = "%s&api_key=%s&artist=%s&track=%s"
	trackSearchURLTempl = "%s&api_key=%s"
)

func CheckAlbumSearchParams(params url.Values) error {
	if _, ok := params["album"]; !ok {
		return domain.RequiredError
	}

	if _, ok := params["artist"]; !ok {
		return domain.RequiredError
	}

	return nil
}

func CheckTrackSearchParams(params url.Values) error {
	if _, ok := params["track"]; !ok {
		return domain.RequiredError
	}

	return nil
}

func GetAlbumByTitleAndArtist(params map[string][]string) (domain.Album, error) {
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

	albumTracks, err := enrichAlbumInfo(album)
	if err != nil {
		return domain.Album{}, fmt.Errorf("enrich: %v", err)
	}

	return albumTracks.Album, nil
}

func enrichAlbumInfo(albumSearch domain.AlbumSearch) (domain.AlbumSearch, error) {
	g := new(errgroup.Group)

	for i := range albumSearch.Album.Tracks {
		track := &albumSearch.Album.Tracks[i]

		g.Go(func() error {
			var trackInfo domain.AlbumTrackSearch

			respBytes, err := getTrackInfo(albumSearch.Artist, track.Name)
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

	if err := g.Wait(); err != nil {
		return domain.AlbumSearch{}, err
	}

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

func GetArtistsAndTracks(params map[string][]string) ([]domain.Track, error) {
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

	tracks, err := enrichTrackInfo(searchPage.Results.TrackMatches.Tracks)
	if err != nil {
		return nil, fmt.Errorf("enrich: %v", err)
	}

	return tracks, nil
}

func enrichTrackInfo(tracks []domain.Track) ([]domain.Track, error) {
	g := new(errgroup.Group)

	for i := range tracks {
		track := &tracks[i]

		g.Go(func() error {
			var trackInfo domain.TrackInfoSearch

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

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return tracks, nil
}
