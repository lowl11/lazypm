package models

import "github.com/lowl11/lazy-collection/array"

type ProjectConfig struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	// database
	UseDatabase bool            `json:"use_database"`
	Database    *DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
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
