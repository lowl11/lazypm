package events

import "lazypm/src/events/skeleton_event"

type ApiEvents struct {
	Skeleton *skeleton_event.Event
}

func Get(basePath string) (*ApiEvents, error) {
	skeleton, err := skeleton_event.Create(basePath)
	if err != nil {
		return nil, err
	}

	return &ApiEvents{
		Skeleton: skeleton,
	}, nil
}
