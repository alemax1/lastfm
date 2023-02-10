package domain

type AlbumSearch struct {
	Album Album `json:"album"`
}

type Album struct {
	Artist    string      `json:"artist"`
	Tags      Tags        `json:"tags"`
	Playcount string      `json:"playcount"`
	Tracks    albumTracks `json:"tracks"`
	Name      string      `json:"name"`
	Listeners string      `json:"listeners"`
}

type albumTracks struct {
	Tracks []albumTrack `json:"track"`
}

type AlbumTrackSearch struct {
	Track albumTrack `json:"track"`
}

type albumTrack struct {
	Name      string `json:"name"`
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
	Tags      Tags   `json:"toptags"`
}

type AlbumSearchRepository interface {
	GetAlbumInfoByTitleAndArtist(params map[string][]string) (*Album, error)
}

type AlbumSearchUsecase interface {
	GetAlbumInfoByTitleAndArtist(params map[string][]string) (*Album, error)
}
