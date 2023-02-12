package router

import (
	handler "api/internal/delivery/http"
	"api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func CreateRoutes(artistUsecase usecase.ArtistTrackSearch, albumUsecase usecase.AlbumSearch) *echo.Echo {
	e := echo.New()

	artistHandler := handler.NewArtistTrackSearch(e, artistUsecase)

	albumHandler := handler.NewAlbumSearch(e, albumUsecase)

	e.GET("/album", albumHandler.GetAlbumInfoByTitleAndArtist)

	e.GET("/tracks", artistHandler.GetTracksAndArtists)

	return e
}
