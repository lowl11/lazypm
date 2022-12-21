package models

type Package struct {
	Project struct {
		Name string `json:"name"`
	} `json:"project"`
}
