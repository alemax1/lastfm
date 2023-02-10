package app

import (
	"api/config"
	"api/internal/delivery/http"
	"api/internal/repository"
	"api/internal/usecase"
	"log"

	"github.com/labstack/echo/v4"
)

func Run() {
	if err := config.Init(); err != nil {
		log.Fatalf("error trying init config: %v", err)
	}

	e := echo.New()

	artistTrackSearchRepository := repository.NewArtistTrackSearchRepository()

	artistTrackSearchUsecase := usecase.NewArtistTrackSearchUsecase(artistTrackSearchRepository)

	albumSearchRepository := repository.NewAlbumSearchRepository()

	albumSearchUsecase := usecase.NewAlbumSearchUsecase(albumSearchRepository)

	http.NewArtistTrackSearchHandler(e, artistTrackSearchUsecase)

	http.NewAlbumSearchHandler(e, albumSearchUsecase)

	e.Logger.Fatal(e.Start("localhost:8088"))
}
