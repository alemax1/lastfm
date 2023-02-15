package app

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"track/config"
	"track/internal/delivery/router"
	"track/internal/repository"
	"track/internal/usecase"

	_ "github.com/lib/pq"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Run() {
	if err := config.Init(); err != nil {
		log.Fatal().Err(err).Msg("error trying init config")
	}

	conn, err := CreateDBConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("error trying create db connection")
	}

	trackPGRepository := repository.NewTrackPG(conn)

	trackUsecase := usecase.NewTrack(trackPGRepository)

	albumPGRepository := repository.NewAlbumPG(conn)

	albumUsecase := usecase.NewAlbumSearch(albumPGRepository)

	e := router.CreateRoutes(trackUsecase, albumUsecase)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))))
}

func CreateDBConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("trackDB.host"),
		viper.GetInt("trackDB.port"),
		viper.GetString("trackDB.user"),
		viper.GetString("trackDB.dbname"),
		viper.GetString("trackDB.password"),
		viper.GetString("trackDB.sslmode"))

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = conn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping: %v", err)
	}

	return conn, nil
}
