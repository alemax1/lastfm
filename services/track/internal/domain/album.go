package domain

type Album struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ArtistID int    `json:"artist_id"`
}

type AlbumResponse struct {
	Data albumData `json:"data"`
}

type albumData struct {
	Artist      string `json:"artist"`
	Name        string `json:"name"`
	Tags        `json:"tags"`
	Playcount   string `json:"playcount"`
	Listeners   string `json:"listeners"`
	albumTracks `json:"tracks"`
}

type albumTracks struct {
	Tracks []albumTrack `json:"track"`
}

type albumTrack struct {
	Name      string `json:"name"`
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
	Tags      `json:"toptags"`
}
