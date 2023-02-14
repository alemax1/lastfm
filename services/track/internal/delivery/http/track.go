package http

import (
	"net/http"
	"strconv"
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

	if _, ok := params["track"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required param"})
	}

	tracks, err := t.track.TrackSearch(params)
	if err != nil {
		log.Err(err).Msg("error calling trackSearch")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, tracks)
}

func (t track) GetTracksByTag(c echo.Context) error {
	params := c.QueryParams()

	if _, ok := params["tag"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required param"})
	}

	page := params.Get("page")

	var pageNum int
	var err error

	if page != "" {
		pageNum, err = strconv.Atoi(page)
		if err != nil || pageNum < 0 {
			return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalidd page"})
		}
	}

	limit := params.Get("limit")

	limitNum := 100

	if limit != "" {
		limitNum, err = strconv.Atoi(limit)
		if err != nil || limitNum < 0 {
			return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid limit"})
		}
	}

	tag := params.Get("tag")

	if tag == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid tag"})
	}

	tracks, err := t.track.GetTracksByTag(pageNum, limitNum, tag)
	if err != nil {
		log.Err(err).Msg("error calling getTracksByTag")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, model.TracksResponse{Data: tracks})
}

func (t track) GetTracksByArtist(c echo.Context) error {
	params := c.QueryParams()

	if _, ok := params["artist"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required param"})
	}

	page := params.Get("page")

	var pageNum int
	var err error

	if page != "" {
		pageNum, err = strconv.Atoi(page)
		if err != nil || pageNum < 0 {
			return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalidd page"})
		}
	}

	limit := params.Get("limit")

	limitNum := 100

	if limit != "" {
		limitNum, err = strconv.Atoi(limit)
		if err != nil || limitNum < 0 {
			return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid limit"})
		}
	}

	artist := params.Get("artist")

	if artist == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid artist"})
	}

	tracks, err := t.track.GetTracksByArtist(pageNum, limitNum, artist)
	if err != nil {
		log.Err(err).Msg("error calling getTracksByArtist")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, model.TracksResponse{Data: tracks})
}
