package http

import (
	"api/internal/delivery/helper"
	"api/internal/delivery/model"
	"api/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type albumSearchHandler struct {
	aUsecase usecase.AlbumSearch
}

func NewAlbumSearch(a usecase.AlbumSearch) *albumSearchHandler {
	return &albumSearchHandler{
		aUsecase: a,
	}
}

func (a albumSearchHandler) GetAlbumInfoByTitleAndArtist(c echo.Context) error {
	var albumSearchParams model.AlbumSearchParams

	if err := c.Bind(&albumSearchParams); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid parameters"})
	}

	if err := helper.CheckAlbumSearchParams(albumSearchParams); err != nil {
		log.Printf("FSDFDDFSDFSDFSFDSFDSFDSDFSDSFSDF")
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	params := c.QueryParams()

	if _, ok := params["album"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required parameter"})
	}

	if _, ok := params["artist"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required parameter"})
	}

	album, err := a.aUsecase.GetAlbumByTitleAndArtist(params)
	if err != nil {
		log.Error().Err(err).Msg("error calling getAlbumByTitleAndArtist")

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, model.AlbumSearchResponse{Data: album})
}
