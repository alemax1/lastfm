package repository

import (
	"api/internal/domain"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type artistTrackSearchRepository struct {
}

func NewArtistTrackSearchRepository() domain.ArtistTrackSearchRepository {
	return new(artistTrackSearchRepository)
}

func (a artistTrackSearchRepository) GetAllArtistsAndTracks() ([]domain.Track, error) {
	resp, err := http.Get("http://ws.audioscrobbler.com/2.0/?method=track.search&track=Ass+like+that&api_key=6741dc681b43487ae782af2c750c7f9d&format=json&artist=Eminem")
	if err != nil {
		return nil, errors.Wrap(err, "get")
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "readAll")
	}
	defer resp.Body.Close()

	var result domain.ArtistTrackSearch

	if err := json.Unmarshal(respBytes, &result); err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	return result.Results.TrackMatches.Tracks, nil
}

func (a artistTrackSearchRepository) GetArtistsAndTracksByPageAndLimit(limit, page int) ([]domain.Track, error) {
	return nil, nil
}
