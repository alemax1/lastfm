package domain

type ArtistTrackSearch struct {
	Results result `json:"results"`
}

type result struct {
	TrackMatches trackMatches `json:"trackmatches"`
}

type trackMatches struct {
	Tracks []Track `json:"track"`
}

type Track struct {
	Name      string `json:"name"`
	Artist    string `json:"artist"`
	Listeners string `json:"listeners"`
}

type ArtistTrackSearchRepository interface {
	GetAllArtistsAndTracks() ([]Track, error)
	GetArtistsAndTracksByPageAndLimit(limit, page int) ([]Track, error)
}

type ArtistTrackSearchUsecase interface {
	GetAllArtistsAndTracks() ([]Track, error)
	GetArtistsAndTracksByPageAndLimit(limit, page int) ([]Track, error)
}
