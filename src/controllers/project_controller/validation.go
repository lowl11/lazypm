package project_controller

import (
	"errors"
	"strings"
)

func (controller *Controller) validateProjectName(name string) error {
	if name == "" {
		return errors.New("project name cannot be empty")
	}

	if strings.Contains(name, " ") {
		return errors.New("project name cannot contain spaces")
	}

	return nil
}

func (controller *Controller) validateProjectDescription(description string) error {
	if description == "" {
		return errors.New("description cannot be empty")
	}

	return nil
}
