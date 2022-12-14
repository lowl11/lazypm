package skeleton_event

import (
	"encoding/json"
	"github.com/lowl11/lazy-collection/array"
	"github.com/lowl11/lazyfile/fileapi"
	"github.com/lowl11/lazyfile/folderapi"
	"lazypm/src/data/entities"
	"lazypm/src/data/models"
	"lazypm/src/services/array_converter"
	"lazypm/src/services/templates"
	"strings"
)

var (
	validateDatabase = array.NewWithList[string]("repositories", "scripts", "start")
)

func (event *Event) validateObject(config *models.ProjectConfig, object *models.SkeletonObject) bool {
	if object.IsFolder {
		if validateDatabase.Contains(object.Name) {
			return config.UseDatabase
		}
	}

	return true
}

func (event *Event) createObject(object *models.SkeletonObject, config *models.ProjectConfig) error {
	// create object if it does not exist
	if object.IsFolder && !object.Exist && event.validateObject(config, object) {
		if err := folderapi.Create(object.Path, object.Name); err != nil {
			return err
		}
	}

	// create children folders recursively
	if object.Children != nil && object.Children.Size() > 0 {
		for _, child := range object.Children.Slice() {
			if child.IsFolder {
				if err := event.createObject(&child, config); err != nil {
					return err
				}
			} else {
				if err := event.createFile(&child, config); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (event *Event) createFile(file *models.SkeletonObject, config *models.ProjectConfig) error {
	// prepare variables
	event.variables["project_name"] = config.Name
	event.variables["project_description"] = config.Description

	if config.UseDatabase {
		event.variables["config_database"] = templates.ConfigDatabase
		event.variables["definition_config_database"] = templates.DefinitionConfigDatabase
	}

	// create file if it does not exist
	if !file.Exist {
		fileTemplate := templates.Get(file.Template)

		// replace variables
		for key, value := range event.variables {
			variableKey := "<% " + key + " %>"
			if strings.Contains(fileTemplate, variableKey) {
				fileTemplate = strings.ReplaceAll(fileTemplate, variableKey, value)
			}
		}

		fileContent := []byte(fileTemplate)
		if err := fileapi.Create(file.Path+"/"+file.Name, fileContent); err != nil {
			return err
		}
	}

	return nil
}

func (event *Event) loadTemplate() error {
	// read template
	templateInBytes, err := fileapi.Read("template.json")
	if err != nil {
		return err
	}

	// parse template
	template := make([]entities.SkeletonObject, 0)
	if err = json.Unmarshal(templateInBytes, &template); err != nil {
		return err
	}

	event.template = array_converter.ConvertSkeletonObjects(event.basePath, template)
	return nil
}

func (event *Event) loadStructure() error {
	// build real structure

	defer event.template.Reset()
	for event.template.Next() {
		obj := event.template.Value()

		// build path
		objectPath := obj.Path
		if objectPath == "/" {
			objectPath = ""
		}
		path := event.basePath
		pathWithName := path + obj.Name

		// runtime fields
		var exist bool
		var empty bool

		if obj.IsFolder {
			exist = folderapi.Exist(pathWithName)
			empty = folderapi.Empty(pathWithName)
		} else {
			exist = fileapi.Exist(pathWithName)
			empty = fileapi.Empty(pathWithName)
		}

		event.skeleton.Push(models.SkeletonObject{
			Name:     obj.Name,
			IsFolder: obj.IsFolder,
			Path:     path,

			// file/folder fields
			Template: obj.Template,
			Children: obj.Children,

			// runtime fields
			Exist: exist,
			Empty: empty,
		})
	}

	return nil
}

func (event *Event) loadVariables() {
	event.variables["project_name"] = "project_name"
	event.variables["project_description"] = "Your Project Description"
	event.variables["port"] = "8080"

	event.variables["config_database"] = ""
	event.variables["definition_config_database"] = ""
}
