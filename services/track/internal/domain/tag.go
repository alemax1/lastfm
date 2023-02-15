package domain

type tag struct {
	Name string `json:"name"`
}

type Tags struct {
	Tags []tag `json:"tag"`
}
