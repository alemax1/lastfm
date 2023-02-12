package app

import (
	"api/config"
	"api/internal/repository"
	"api/internal/usecase"
	"api/router"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Run() {
	if err := config.Init(); err != nil {
		log.Fatal().Err(err).Msg("error trying init config")
	}

	artistTrackSearchRepository := repository.NewArtistTrackSearch()

	artistTrackSearchUsecase := usecase.NewArtistTrackSearch(artistTrackSearchRepository)

	albumSearchRepository := repository.NewAlbumSearch()

	albumSearchUsecase := usecase.NewAlbumSearch(albumSearchRepository)

	e := router.CreateRoutes(artistTrackSearchUsecase, albumSearchUsecase)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))))
}
