package skeleton_event

import (
	"github.com/lowl11/lazy-collection/array"
	"lazypm/src/data/models"
	"lazypm/src/events/pkg_event"
)

type Event struct {
	basePath string
	pkg      *pkg_event.Event

	skeleton *array.Array[models.SkeletonObject]
	template *array.Array[models.SkeletonObject]

	variables map[string]string
}

func Create(basePath string, pkg *pkg_event.Event) (*Event, error) {
	event := &Event{
		basePath: basePath,
		pkg:      pkg,

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
