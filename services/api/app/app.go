package app

import (
	"api/internal/delivery/http"
	"api/internal/repository"
	"api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	artistTrackSearchRepository := repository.NewArtistTrackSearchRepository()

	artistTrackSearchUsecase := usecase.NewArtistTrackSearchUsecase(artistTrackSearchRepository)

	http.NewArtistTrackSearchHandler(e, artistTrackSearchUsecase)

	e.Logger.Fatal(e.Start("localhost:8088"))
}
