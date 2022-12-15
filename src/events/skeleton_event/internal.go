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
	validateDatabase = array.NewWithList[string](
		"repositories", "scripts", "start", "script_event",
	)
)

func (event *Event) validateObject(config *models.ProjectConfig, object *models.SkeletonObject) bool {
	// database case validation
	if object.IsFolder {
		pathContains := validateDatabase.Single(func(item string) bool {
			return strings.Contains(object.Path, item)
		}) != nil

		if validateDatabase.Contains(object.Name) || pathContains {
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
	// if file path contains database path
	if !config.UseDatabase {
		if validateDatabase.Single(func(item string) bool {
			return strings.Contains(file.Path, item)
		}) != nil {
			return nil
		}
	}

	// prepare variables
	event.variables["project_name"] = config.Name
	event.variables["project_description"] = config.Description

	// fill database variables
	if config.UseDatabase {
		// config variables
		event.variables["database_server"] = config.Database.Server
		event.variables["database_port"] = config.Database.Port
		event.variables["database_user"] = config.Database.Username
		event.variables["database_password"] = config.Database.Password
		event.variables["database_name"] = config.Database.Name
		event.variables["database_max_connections"] = config.Database.MaxConnections
		event.variables["database_max_lifetime"] = config.Database.MaxLifetime

		// go mod libraries
		event.variables["gomod_sqlx"] = "github.com/jmoiron/sqlx v1.3.5"
		event.variables["gomod_postgres"] = "github.com/lib/pq v1.2.0"

		// imports
		event.variables["events_import_script"] = "\"" + config.Name + "/src/events/script_event\""
		event.variables["events_field_script"] = "Script *script_event.Event"
		event.variables["events_create_script"] = templates.Get("events_script_create")
		event.variables["events_script_contain"] = templates.Get("events_script_contain")
		event.variables["controllers_import_repositories"] = templates.FillVariables("controllers_import_repositories", event.variables)
		event.variables["controllers_argument_repositories"] = templates.Get("controllers_argument_repositories")
		event.variables["api_import_repositories"] = templates.FillVariables("api_import_repositories", event.variables)
		event.variables["api_repositories_create"] = templates.Get("api_repositories_create")
		event.variables["api_field_repositories"] = templates.Get("api_field_repositories")

		// fill templates
		event.variables["config_database"] = templates.FillVariables("config_database", event.variables)
		event.variables["definition_config_database"] = templates.FillVariables("definition_config_database", event.variables)
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

	// database variables
	event.variables["database_server"] = ""
	event.variables["database_port"] = ""
	event.variables["database_user"] = ""
	event.variables["database_password"] = ""
	event.variables["database_name"] = ""
	event.variables["database_max_connections"] = ""
	event.variables["database_max_lifetime"] = ""
	event.variables["gomod_sqlx"] = ""
	event.variables["gomod_postgres"] = ""

	// fill templates
	event.variables["config_database"] = ""
	event.variables["definition_config_database"] = ""

	// database imports
	event.variables["events_import_script"] = ""
	event.variables["events_field_script"] = ""
	event.variables["events_create_script"] = ""
	event.variables["events_script_contain"] = "//"
	event.variables["controllers_import_repositories"] = ""
	event.variables["controllers_argument_repositories"] = ""
	event.variables["api_import_repositories"] = ""
	event.variables["api_repositories_create"] = ""
	event.variables["api_field_repositories"] = ""
}
