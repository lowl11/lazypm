package controllers

import (
	"lazypm/src/controllers/controller_controller"
	"lazypm/src/controllers/project_controller"
	"lazypm/src/controllers/repository_controller"
	"lazypm/src/events"
)

type ApiControllers struct {
	Project    *project_controller.Controller
	Controller *controller_controller.Controller
	Repository *repository_controller.Controller
}

func Get(apiEvents *events.ApiEvents) *ApiControllers {
	return &ApiControllers{
		Project:    project_controller.Create(apiEvents.Skeleton),
		Controller: controller_controller.Create(apiEvents.Skeleton),
		Repository: repository_controller.Create(),
	}
}
