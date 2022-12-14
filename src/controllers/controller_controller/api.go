package controller_controller

import (
	"github.com/lowl11/lazy-cli/cli_route"
	"lazypm/src/data/errors"
	"lazypm/src/definition"
)

func (controller *Controller) New(ctx cli_route.IContext) error {
	printer := definition.Printer
	params := ctx.Params()
	if len(params) == 0 {
		return errors.ControllerNewParam
	}

	controllerName := params[0]

	printer.Info("Create new controller: " + controllerName + "...")
	return nil
}
