package domain

type ArtistTrackSearch struct {
	Results `json:"results"`
}

type Results struct {
	TotalResult  string `json:"opensearch:totalResults"`
	ItemsPerPage string `json:"opensearch:itemsPerPage"`
	TrackMatches `json:"trackmatches"`
}

type TrackInfoSearch struct {
	TrackInfo trackInfo `json:"track"`
}

type trackInfo struct {
	Playcount string `json:"playcount"`
	TopTags   Tags   `json:"toptags"`
}

type Track struct {
	Name      string `json:"name"`
	Artist    string `json:"artist"`
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
	Tags      []Tag  `json:"tags"`
}

type TrackMatches struct {
	Tracks []Track `json:"track"`
}

type ArtistTrackSearchRepository interface {
	GetArtistsAndTracksByPageAndLimit(params map[string][]string) ([]Track, error)
}

type ArtistTrackSearchUsecase interface {
	GetArtistsAndTracksByPageAndLimit(params map[string][]string) ([]Track, error)
}
