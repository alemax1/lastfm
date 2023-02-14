package domain

type AlbumSearch struct {
	Album `json:"album"`
}

type Album struct {
	Artist      string `json:"artist"`
	Name        string `json:"name"`
	Tags        `json:"tags"`
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

type AlbumTrackSearch struct {
	AlbumTrack `json:"track"`
}
