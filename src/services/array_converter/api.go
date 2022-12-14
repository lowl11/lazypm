package array_converter

import (
	"github.com/lowl11/lazy-collection/array"
	"lazypm/src/data/entities"
	"lazypm/src/data/models"
)

func ConvertSkeletonObjects(basePath string, list []entities.SkeletonObject) *array.Array[models.SkeletonObject] {
	newArray := array.NewWithSize[models.SkeletonObject](len(list))

	for _, item := range list {
		newPath := basePath + item.Path

		newItem := models.SkeletonObject{
			Name:     item.Name,
			IsFolder: item.IsFolder,
			Path:     newPath,
			Template: item.Template,
		}

		if item.IsFolder && len(item.Children) > 0 {
			newItem.Children = ConvertSkeletonObjects(basePath, item.Children)
		}

		newArray.Push(newItem)
	}

	return newArray
}
