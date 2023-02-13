package http

import (
	"net/http"
	"track/internal/delivery/model"
	"track/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type albumSearch struct {
	albumSearch usecase.AlbumSearch
}

func NewAlbumSearch(e *echo.Echo, a usecase.AlbumSearch) *albumSearch {
	return &albumSearch{
		albumSearch: a,
	}
}

func (a albumSearch) AlbumSearch(c echo.Context) error {
	params := c.QueryParams()

	if _, ok := params["album"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required parameter"})
	}

	if _, ok := params["artist"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required parameter"})
	}

	album, err := a.albumSearch.AlbumSearch(params)
	if err != nil {
		log.Err(err).Msg("error calling albumSearch")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, album)
}
