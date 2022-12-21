package models

type Package struct {
	Project struct {
		Name       string `json:"name"`
		IsDatabase bool   `json:"is_database"`
	} `json:"project"`
}
