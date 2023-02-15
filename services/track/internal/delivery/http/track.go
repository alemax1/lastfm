package http

import (
	"net/http"
	"track/internal/delivery/helper"
	"track/internal/delivery/model"
	"track/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type track struct {
	track usecase.Track
}

func NewTrackSearch(e *echo.Echo, t usecase.Track) *track {
	return &track{
		track: t,
	}
}

func (t track) TrackSearch(c echo.Context) error {
	params := c.QueryParams()

	if err := helper.CheckTrackSearchParams(params); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	tracks, err := helper.TrackSearch(params)
	if err != nil {
		log.Err(err).Msg("error calling trackSearch helper")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	if err := t.track.TrackSearch(c.Request().Context(), params, tracks); err != nil {
		log.Err(err).Msg("error calling trackSearch")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, tracks)
}

func (t track) GetTracksByTag(c echo.Context) error {
	params := c.QueryParams()

	tag, err := helper.CheckTracksByTagParams(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	page, limit, err := helper.GetPagination(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid pagination params"})
	}

	tracks, err := t.track.GetTracksByTag(c.Request().Context(), page, limit, tag)
	if err != nil {
		log.Err(err).Msg("error calling getTracksByTag")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, model.TracksResponse{Data: tracks})
}

func (t track) GetTracksByArtist(c echo.Context) error {
	params := c.QueryParams()

	artist, err := helper.CheckTracksByArtistParams(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	page, limit, err := helper.GetPagination(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid pagination params"})
	}

	tracks, err := t.track.GetTracksByArtist(c.Request().Context(), page, limit, artist)
	if err != nil {
		log.Err(err).Msg("error calling getTracksByArtist")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, model.TracksResponse{Data: tracks})
}
