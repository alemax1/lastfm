package domain

type Track struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Listeners  int    `json:"listeners"`
	Playcounts int    `json:"playcounts"`
	ArtistID   int    `json:"artist_id"`
}

type TrackDBResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Listeners  int    `json:"listeners"`
	Playcounts int    `json:"playcounts"`
	Artist     string `json:"artist"`
}

type TrackResponse struct {
	Name      string `json:"name"`
	Artist    string `json:"artist"`
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
	Tags      []Tag  `json:"tags"`
}

type TracksResponse struct {
	Data []TrackResponse `json:"data"`
}
