package http

import (
	"api/internal/delivery/helper"
	"api/internal/delivery/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type albumSearchHandler struct {
}

func NewAlbumSearch() *albumSearchHandler {
	return new(albumSearchHandler)
}

func (a albumSearchHandler) GetAlbumInfoByTitleAndArtist(c echo.Context) error {
	params := c.QueryParams()

	if err := helper.CheckAlbumSearchParams(params); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	album, err := helper.GetAlbumByTitleAndArtist(params)
	if err != nil {
		log.Error().Err(err).Msg("error calling getAlbumByTitleAndArtist")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, model.AlbumSearchResponse{Data: album})
}
