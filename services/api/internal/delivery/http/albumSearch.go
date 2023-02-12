package http

import (
	"api/internal/delivery/model"
	"api/internal/usecase"
	"net/http"

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

func (a albumSearch) GetAlbumInfoByTitleAndArtist(c echo.Context) error {
	params := c.QueryParams()

	if _, ok := params["album"]; !ok { //mapper
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required parameter"})
	}

	if _, ok := params["artist"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required parameter"})
	}

	album, err := a.albumSearch.GetAlbumByTitleAndArtist(params)
	if err != nil {
		log.Error().Err(err).Msg("error calling getAlbumByTitleAndArtist")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, model.AlbumSearchResponse{Data: album})
}
