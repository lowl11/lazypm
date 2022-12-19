package controller_controller

import (
	"github.com/lowl11/lazy-cli/cli_route"
	"lazypm/src/data/models"
	"lazypm/src/definition"
	"lazypm/src/services/cmd_tool"
)

func (controller *Controller) New(ctx cli_route.IContext) error {
	printer := definition.Printer
	//params := ctx.Params()
	//if len(params) == 0 {
	//	return errors.ControllerNewParam
	//}

	controllerName := cmd_tool.Ask("Controller name")
	if err := controller.validateControllerName(controllerName); err != nil {
		printer.Error(err)
		controllerName = cmd_tool.AskAlways("Controller name", controller.validateControllerName)
	}

	printer.Info("Create new controller: " + controllerName + "...")

	if err := controller.skeleton.NewController(&models.ControllerConfig{
		Name: controllerName,
	}); err != nil {
		return err
	}

	return nil
}
