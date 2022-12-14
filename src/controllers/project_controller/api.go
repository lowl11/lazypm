package project_controller

import (
	"github.com/lowl11/lazy-cli/cli_route"
	"lazypm/src/data/errors"
	"lazypm/src/data/models"
	"lazypm/src/definition"
)

func (controller *Controller) Init(ctx cli_route.IContext) error {
	printer := definition.Printer
	params := ctx.Params()
	if len(params) == 0 {
		return errors.ProjectInitParam
	}

	projectName := params[0]

	printer.Info("Create project: " + projectName + "...")

	// create not exist folders & files
	config := &models.ProjectConfig{
		Name: projectName,
	}
	if err := controller.skeleton.CreateObjects(config); err != nil {
		return err
	}

	return nil
}
