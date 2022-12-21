package skeleton_event

import (
	"errors"
	"github.com/lowl11/lazy-collection/array"
	"lazypm/src/data/models"
)

func (event *Event) Variables(variables map[string]string) {
	for key, value := range variables {
		event.variables[key] = value
	}
}

func (event *Event) CreateObjects(config *models.ProjectConfig) error {
	defer event.skeleton.Reset()
	for event.skeleton.Next() {
		obj := event.skeleton.Value()

		if obj.IsFolder {
			if err := event.createObject(&obj, config); err != nil {
				return err
			}
		} else {
			if err := event.createFile(&obj, config); err != nil {
				return err
			}
		}
	}

	return nil
}

func (event *Event) NewController(config *models.ControllerConfig) error {
	// package event check
	if event.pkg.Get() == nil {
		return errors.New("package config is empty or null")
	}

	projectInfo := event.pkg.Get()

	// get /src folders
	src := event.skeleton.Single(func(item models.SkeletonObject) bool {
		return item.Name == "src"
	})
	if src == nil {
		return errors.New("src folder not found")
	}

	// get /src/controllers folder
	controllers := src.Children.Single(func(item models.SkeletonObject) bool {
		return item.Name == "controllers"
	})
	if controllers == nil {
		return errors.New("controllers folder not found")
	}

	// create new "name"_controller folder
	controllerFolder := config.Name + "_controller"

	controllers.Children.Push(models.SkeletonObject{
		Name:     controllerFolder,
		IsFolder: true,
		Path:     event.basePath + "/src/controllers",
		Children: array.NewWithList[models.SkeletonObject](
			models.SkeletonObject{
				Name:     "api.go",
				IsFolder: false,
				Path:     event.basePath + "/src/controllers/" + controllerFolder,
				Template: "controllers_new_api",
			},
			models.SkeletonObject{
				Name:     "controller.go",
				IsFolder: false,
				Path:     event.basePath + "/src/controllers/" + controllerFolder,
				Template: "controllers_new_controller",
			},
			models.SkeletonObject{
				Name:     "validation.go",
				IsFolder: false,
				Path:     event.basePath + "/src/controllers/" + controllerFolder,
				Template: "controllers_new_validation",
			},
		),
	})

	if err := event.createObject(controllers, &models.ProjectConfig{
		Name:        projectInfo.Project.Name,
		UseDatabase: projectInfo.Project.IsDatabase,
	}); err != nil {
		return err
	}

	return nil
}
