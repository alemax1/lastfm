package domain

type ArtistTrackSearch struct {
	Results `json:"results"`
}

type Results struct {
	TotalResult  string `json:"opensearch:totalResults"`
	ItemsPerPage string `json:"opensearch:itemsPerPage"`
	TrackMatches `json:"trackmatches"`
}

type TrackMatches struct {
	Tracks []Track `json:"track"`
}

type TrackInfoSearch struct {
	TrackInfo `json:"track"`
}

type TrackInfo struct {
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
