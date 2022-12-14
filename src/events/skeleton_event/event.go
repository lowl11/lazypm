package skeleton_event

import (
	"github.com/lowl11/lazy-collection/array"
	"lazypm/src/data/models"
)

type Event struct {
	basePath string

	skeleton *array.Array[models.SkeletonObject]
	template *array.Array[models.SkeletonObject]

	variables map[string]string
}

func Create(basePath string) (*Event, error) {
	event := &Event{
		basePath:  basePath,
		skeleton:  array.New[models.SkeletonObject](),
		variables: make(map[string]string),
	}

	// load template
	if err := event.loadTemplate(); err != nil {
		return nil, err
	}

	// load structure
	if err := event.loadStructure(); err != nil {
		return nil, err
	}

	event.loadVariables()

	return event, nil
}
