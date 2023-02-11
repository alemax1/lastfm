package domain

type AlbumSearch struct {
	Album `json:"album"`
}

type Album struct {
	Artist      string `json:"artist"`
	Name        string `json:"name"`
	Tags        Tags   `json:"tags"`
	Playcount   string `json:"playcount"`
	albumTracks `json:"tracks"`
	Listeners   string `json:"listeners"`
}

type albumTracks struct {
	Tracks []AlbumTrack `json:"track"`
}

type AlbumTrackSearch struct {
	Track AlbumTrack `json:"track"`
}

type AlbumTrack struct {
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
