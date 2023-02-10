package http

import (
	"api/internal/delivery/model"
	"api/internal/domain"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type albumSearchHandler struct {
	albumSearchUsecase domain.AlbumSearchUsecase
}

func NewAlbumSearchHandler(e *echo.Echo, a domain.AlbumSearchUsecase) {
	handler := &albumSearchHandler{
		albumSearchUsecase: a,
	}

	e.GET("/album", handler.getAlbumInfoByTitleAndArtist)
}

func (a albumSearchHandler) getAlbumInfoByTitleAndArtist(c echo.Context) error {
	params := c.QueryParams()

	if _, ok := params["album"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required parameter"})
	}

	if _, ok := params["artist"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required parameter"})
	}

	album, err := a.albumSearchUsecase.GetAlbumInfoByTitleAndArtist(params)
	if err != nil {
		log.Printf("error calling getAlbumInfoByTitleAndArtist: %v", err)

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, model.AlbumSearchResponse{Data: *album})
}
