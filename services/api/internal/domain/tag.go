package domain

type Tag struct {
	Name string `json:"name"`
}

type Tags struct {
	Tag []Tag `json:"tag"`
}
