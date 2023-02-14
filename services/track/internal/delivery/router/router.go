package router

import (
	"github.com/labstack/echo/v4"

	handler "track/internal/delivery/http"
	"track/internal/usecase"
)

func CreateRoutes(trackSearchUsecase usecase.Track, albumSearchUsecase usecase.AlbumSearch) *echo.Echo {
	e := echo.New()

	trackHandler := handler.NewTrackSearch(e, trackSearchUsecase)

	albumHandler := handler.NewAlbumSearch(e, albumSearchUsecase)

	e.GET("/tracks", trackHandler.TrackSearch)

	e.GET("/album", albumHandler.AlbumSearch)

	e.GET("/tags/tracks", trackHandler.GetTracksByTag)

	e.GET("/artists/tracks", trackHandler.GetTracksByArtist)

	return e
}
