package templates

const (
	eventsApi = `package events

import (
	"<% project_name %>/src/events/script_event"
)

type ApiEvents struct {
	Script *script_event.Event
}

func Get() (*ApiEvents, error) {
	scriptEvent, err := script_event.Create()
	if err != nil {
		return nil, err
	}

	return &ApiEvents{
		Script: scriptEvent,
	}, nil
}`

	eventsScriptApi = `package script_event

func (event *Event) StartScript(script string) string {
	return event.startScripts[script+".sql"]
}

func (event *Event) Script(folder, script string) string {
	return event.scripts[folder].(map[string]string)[script+".sql"]
}`

	eventsScriptEvent = `package script_event

type Event struct {
	startScripts map[string]string
	scripts      map[string]any
}

func Create() (*Event, error) {
	event := &Event{
		startScripts: make(map[string]string),
		scripts:      make(map[string]any),
	}

	if err := event.readStartScripts(); err != nil {
		return nil, err
	}

	if err := event.readScripts(); err != nil {
		return nil, err
	}

	return event, nil
}`

	eventsScriptInternal = `package script_event

import (
	"github.com/lowl11/lazyfile/folderapi"
)

func (event *Event) readStartScripts() error {
	files, err := folderapi.Objects("resources/scripts/start")
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsFolder {
			continue
		}

		body, err := file.Read()
		if err != nil {
			return err
		}

		event.startScripts[file.Name] = string(body)
	}

	return nil
}

func (event *Event) readScripts() error {
	folders, err := folderapi.Objects("resources/scripts/")
	if err != nil {
		return err
	}

	for _, folder := range folders {
		if !folder.IsFolder {
			continue
		}

		folderMap := make(map[string]string)

		files, err := folderapi.Objects("resources/scripts/" + folder.Name)
		if err != nil {
			return err
		}

		for _, file := range files {
			body, err := file.Read()
			if err != nil {
				return err
			}

			folderMap[file.Name] = string(body)
		}

		event.scripts[folder.Name] = folderMap
	}

	return nil
}`
)
