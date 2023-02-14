package domain

type Album struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ArtistID int    `json:"artist_id"`
}

type Response struct {
	Data AlbumResponse `json:"data"`
}

type AlbumResponse struct {
	Artist      string `json:"artist"`
	Name        string `json:"name"`
	Tags        `json:"tags"`
	Playcount   string `json:"playcount"`
	Listeners   string `json:"listeners"`
	AlbumTracks `json:"tracks"`
}

type AlbumTracks struct {
	Tracks []AlbumTrack `json:"track"`
}

type AlbumTrack struct {
	Name      string `json:"name"`
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
	Tags      `json:"toptags"`
}
