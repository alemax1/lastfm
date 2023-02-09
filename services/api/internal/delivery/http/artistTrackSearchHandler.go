package http

import (
	"api/internal/delivery/model"
	"api/internal/domain"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type artistTrackSearchHandler struct {
	artistTrackSearchUseCase domain.ArtistTrackSearchUsecase
}

func NewArtistTrackSearchHandler(e *echo.Echo, a domain.ArtistTrackSearchUsecase) {
	handler := &artistTrackSearchHandler{
		artistTrackSearchUseCase: a,
	}

	e.GET("/tracks", handler.GetAllTracksAndArtists)
}

func (a artistTrackSearchHandler) GetAllTracksAndArtists(c echo.Context) error {
	tracks, err := a.artistTrackSearchUseCase.GetAllArtistsAndTracks()
	if err != nil {
		log.Printf("error calling GetAllArtistsAndTracks: %v", err)

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	fmt.Println(len(tracks))

	return c.JSON(http.StatusOK, model.TrackAndArtistsResponse{Data: tracks})
}
