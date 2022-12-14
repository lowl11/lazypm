package project_controller

import (
	"errors"
	"regexp"
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

// database

func (controller *Controller) validateDatabaseServer(server string) error {
	if server == "" {
		return errors.New("database server cannot be empty")
	}

	reg, _ := regexp.Compile("^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}$")
	if !reg.MatchString(server) {
		return errors.New("database server must in format 0.0.0.0")
	}

	return nil
}

func (controller *Controller) validateDatabasePort(port string) error {
	if port == "" {
		return errors.New("database port cannot be empty")
	}

	reg, _ := regexp.Compile("[0-9]+")
	if !reg.MatchString(port) {
		return errors.New("database port must contain only digits")
	}

	return nil
}

func (controller *Controller) validateDatabaseUsername(username string) error {
	if username == "" {
		return errors.New("database username cannot be empty")
	}

	return nil
}

func (controller *Controller) validateDatabasePassword(password string) error {
	if password == "" {
		return errors.New("database password cannot be empty")
	}

	return nil
}

func (controller *Controller) validateDatabaseName(name string) error {
	if name == "" {
		return errors.New("database name cannot be empty")
	}

	return nil
}
