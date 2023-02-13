package http

import (
	"net/http"
	"track/internal/delivery/model"
	"track/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type trackSearch struct {
	trackSearch usecase.TrackSearch
}

func NewTrackSearch(e *echo.Echo, t usecase.TrackSearch) *trackSearch {
	return &trackSearch{
		trackSearch: t,
	}
}

func (t trackSearch) TrackSearch(c echo.Context) error {
	params := c.QueryParams()

	if _, ok := params["track"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required param"})
	}

	tracks, err := t.trackSearch.TrackSearch(params)
	if err != nil {
		log.Err(err).Msg("error calling trackSearch")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, tracks)
}
