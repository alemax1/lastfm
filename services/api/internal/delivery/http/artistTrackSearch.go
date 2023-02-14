package http

import (
	"api/internal/delivery/helper"
	"api/internal/delivery/model"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/labstack/echo/v4"
)

type artistTrackSearchHandler struct {
}

func NewArtistTrackSearch() *artistTrackSearchHandler {
	return new(artistTrackSearchHandler)
}

func (a artistTrackSearchHandler) GetTracksAndArtists(c echo.Context) error {
	params := c.QueryParams()

	if err := helper.CheckTrackSearchParams(params); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	tracks, err := helper.GetArtistsAndTracks(params)
	if err != nil {
		log.Error().Err(err).Msg("error calling getAlbumByTitleAndArtist")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, model.TrackAndArtistsResponse{Data: tracks})
}
