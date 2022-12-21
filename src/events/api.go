package events

import (
	"lazypm/src/events/pkg_event"
	"lazypm/src/events/skeleton_event"
)

type ApiEvents struct {
	Skeleton *skeleton_event.Event
	Package  *pkg_event.Event
}

func Get(basePath string) (*ApiEvents, error) {
	pkg, err := pkg_event.Create(basePath)
	if err != nil {
		return nil, err
	}

	skeleton, err := skeleton_event.Create(basePath, pkg)
	if err != nil {
		return nil, err
	}

	return &ApiEvents{
		Skeleton: skeleton,
		Package:  pkg,
	}, nil
}
