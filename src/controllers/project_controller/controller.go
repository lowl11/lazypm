package project_controller

import "lazypm/src/events/skeleton_event"

type Controller struct {
	skeleton *skeleton_event.Event
}

func Create(skeleton *skeleton_event.Event) *Controller {
	return &Controller{
		skeleton: skeleton,
	}
}
