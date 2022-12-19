package skeleton_event

import (
	"errors"
	"github.com/lowl11/lazy-collection/array"
	"lazypm/src/data/models"
)

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
	// get /src folders
	src := event.skeleton.Single(func(item models.SkeletonObject) bool {
		return item.Name == "src"
	})
	if src != nil {
		return errors.New("src folder not found")
	}

	// get /src/controllers folder
	controllers := event.skeleton.Single(func(item models.SkeletonObject) bool {
		return item.Name == "controllers"
	})
	if controllers == nil {
		return errors.New("controllers folder not found")
	}

	// create new "name"_controller folder
	controllers.Children.Push(models.SkeletonObject{
		Name:     config.Name + "_controller",
		IsFolder: true,
		Path:     "/src/controllers",
		Children: array.NewWithList[models.SkeletonObject](
			models.SkeletonObject{
				Name:     "api.go",
				IsFolder: false,
				Path:     "/src/controllers/" + config.Name,
				Template: "controllers_new_api",
			},
			models.SkeletonObject{
				Name:     "controller.go",
				IsFolder: false,
				Path:     "/src/controllers/" + config.Name,
				Template: "controllers_new_controller",
			},
			models.SkeletonObject{
				Name:     "validation.go",
				IsFolder: false,
				Path:     "/src/controllers/" + config.Name,
				Template: "controllers_new_validation",
			},
		),
	})

	if err := event.createObject(controllers, &models.ProjectConfig{
		Name: "test-service",
	}); err != nil {
		return err
	}

	return nil
}
