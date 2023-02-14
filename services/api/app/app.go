package app

import (
	"api/config"
	"api/internal/delivery/router"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Run() {
	if err := config.Init(); err != nil {
		log.Fatal().Err(err).Msg("error trying init config")
	}

	e := router.CreateRoutes()

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))))
}
