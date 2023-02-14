package http

import (
	"api/internal/delivery/model"
	"api/internal/usecase"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/labstack/echo/v4"
)

type artistTrackSearch struct {
	artistTrackSearch usecase.ArtistTrackSearch
}

func NewArtistTrackSearch(a usecase.ArtistTrackSearch) *artistTrackSearch {
	return &artistTrackSearch{
		artistTrackSearch: a,
	}
}

func (a artistTrackSearch) GetTracksAndArtists(c echo.Context) error {
	params := c.QueryParams()

	if _, ok := params["track"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required parameter"})
	}

	tracks, err := a.artistTrackSearch.GetArtistsAndTracks(params)
	if err != nil {
		log.Error().Err(err).Msg("error calling getAlbumByTitleAndArtist")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, model.TrackAndArtistsResponse{Data: tracks})
}
