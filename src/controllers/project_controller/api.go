package project_controller

import (
	"github.com/lowl11/lazy-cli/cli_route"
	"lazypm/src/data/models"
	"lazypm/src/definition"
	"lazypm/src/services/cmd_tool"
)

func (controller *Controller) Init(ctx cli_route.IContext) error {
	printer := definition.Printer
	//params := ctx.Params()
	//if len(params) == 0 {
	//	return errors.ProjectInitParam
	//}

	// start collecting params
	// project name
	projectName := cmd_tool.Ask("Project Name")
	if err := controller.validateProjectName(projectName); err != nil {
		printer.Error(err)
		projectName = cmd_tool.AskAlways("Project Name", controller.validateProjectName)
	}

	// project description
	projectDescription := cmd_tool.Ask("Project Description")
	if err := controller.validateProjectDescription(projectDescription); err != nil {
		printer.Error(err)
		projectDescription = cmd_tool.AskAlways("Project Description", controller.validateProjectDescription)
	}

	printer.Info("Create project: " + projectName + "...")

	// create not exist folders & files
	config := &models.ProjectConfig{
		Name:        projectName,
		Description: projectDescription,
	}
	if err := controller.skeleton.CreateObjects(config); err != nil {
		return err
	}

	return nil
}
