package http

import (
	"net/http"
	"track/internal/delivery/helper"
	"track/internal/delivery/model"
	"track/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type albumSearch struct {
	aUsecase usecase.AlbumSearch
}

func NewAlbumSearch(e *echo.Echo, a usecase.AlbumSearch) *albumSearch {
	return &albumSearch{
		aUsecase: a,
	}
}

func (a albumSearch) AlbumSearch(c echo.Context) error {
	params := c.QueryParams()

	if err := helper.CheckAlbumSearchParams(params); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	album, err := helper.AlbumSearch(params)
	if err != nil {
		log.Err(err).Msg("error calling albumSearch helper")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	if err := a.aUsecase.AlbumSearch(c.Request().Context(), params, album); err != nil {
		log.Err(err).Msg("error calling albumSearch")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, album)
}
