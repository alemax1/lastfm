package router

import (
	"github.com/labstack/echo/v4"

	handler "track/internal/delivery/http"
	"track/internal/usecase"
)

func CreateRoutes(trackSearchUsecase usecase.TrackSearch, albumSearchUsecase usecase.AlbumSearch) *echo.Echo {
	e := echo.New()

	trackHandler := handler.NewTrackSearch(e, trackSearchUsecase)

	albumHandler := handler.NewAlbumSearch(e, albumSearchUsecase)

	e.GET("/tracks", trackHandler.TrackSearch)

	e.GET("/album", albumHandler.AlbumSearch)

	return e
}
