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
	Tags        Tags   `json:"tags"`
	Playcount   string `json:"playcount"`
	AlbumTracks `json:"tracks"`
	Listeners   string `json:"listeners"`
}

type AlbumTracks struct {
	Tracks []AlbumTrack `json:"track"`
}

type AlbumTrack struct {
	Name      string `json:"name"`
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
	Tags      Tags   `json:"toptags"`
}
