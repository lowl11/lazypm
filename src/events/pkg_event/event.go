package pkg_event

import "lazypm/src/data/models"

type Event struct {
	basePath string

	pkg *models.Package
}

func Create(basePath string) (*Event, error) {
	event := &Event{
		basePath: basePath,
	}

	if err := event.read(); err != nil {
		return nil, err
	}

	return event, nil
}
