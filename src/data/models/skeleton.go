package models

import "github.com/lowl11/lazy-collection/array"

type ProjectConfig struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UseDatabase bool   `json:"use_database"`
}

type SkeletonObject struct {
	Name     string `json:"name"`
	IsFolder bool   `json:"is_folder"`
	Path     string `json:"path"`

	// file/folder fields
	Template string                       `json:"template,omitempty"`
	Children *array.Array[SkeletonObject] `json:"children,omitempty"`

	// runtime options
	Exist bool `json:"exist"`
	Empty bool `json:"empty"`
}
