package http

import (
	"api/internal/delivery/model"
	"api/internal/domain"
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

	e.GET("/tracks", handler.getAllTracksAndArtists)
}

func (a artistTrackSearchHandler) getAllTracksAndArtists(c echo.Context) error {
	params := c.QueryParams()

	if _, ok := params["track"]; !ok {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "missing required parameter"})
	}

	tracks, err := a.artistTrackSearchUseCase.GetArtistsAndTracksByPageAndLimit(params)
	if err != nil {
		log.Printf("error calling GetAllArtistsAndTracks: %v", err)

		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, model.TrackAndArtistsResponse{Data: tracks})
}
