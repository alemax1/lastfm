package domain

type Tags struct {
	Tags []Tag `json:"tag,omitempty"`
}

type Tag struct {
	Name string `json:"name"`
}
