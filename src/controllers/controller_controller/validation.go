package controller_controller

import "errors"

func (controller *Controller) validateControllerName(name string) error {
	if name == "" {
		return errors.New("controller name cannot be empty")
	}

	return nil
}
