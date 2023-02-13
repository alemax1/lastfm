package domain

type Tag struct {
	Name string `json:"name"`
}

type Tags struct {
	Tags []Tag `json:"tag"`
}
