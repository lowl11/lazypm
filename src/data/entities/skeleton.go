package entities

type SkeletonObject struct {
	Name     string `json:"name"`
	IsFolder bool   `json:"is_folder"`
	Path     string `json:"path"`

	// file/folder fields
	Template string           `json:"template,omitempty"`
	Children []SkeletonObject `json:"children,omitempty"`
}
