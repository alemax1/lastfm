package router

import (
	handler "api/internal/delivery/http"

	"github.com/labstack/echo/v4"
)

func CreateRoutes() *echo.Echo {
	e := echo.New()

	artistHandler := handler.NewArtistTrackSearch()

	albumSearchHandler := handler.NewAlbumSearch()

	e.GET("/album", albumSearchHandler.GetAlbumInfoByTitleAndArtist)

	e.GET("/tracks", artistHandler.GetTracksAndArtists)

	return e
}
