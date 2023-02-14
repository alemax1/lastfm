package domain

type Tags struct {
	Tags []Tag `json:"tag"`
}

type Tag struct {
	Name string `json:"name"`
}
