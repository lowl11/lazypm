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

	// use database
	useDatabase := cmd_tool.Confirm("Use Database?")
	var databaseServer string
	var databasePort string
	var databaseUsername string
	var databasePassword string
	var databaseName string

	if useDatabase {
		databaseServer = cmd_tool.Ask("Database server")
		if err := controller.validateDatabaseServer(databaseServer); err != nil {
			printer.Error(err)
			databaseServer = cmd_tool.AskAlways("Database server", controller.validateDatabaseServer)
		}

		databasePort = cmd_tool.Ask("Database port")
		if err := controller.validateDatabasePort(databasePort); err != nil {
			printer.Error(err)
			databasePort = cmd_tool.AskAlways("Database port", controller.validateDatabasePort)
		}

		databaseUsername = cmd_tool.Ask("Database username")
		if err := controller.validateDatabaseUsername(databaseUsername); err != nil {
			printer.Error(err)
			databaseUsername = cmd_tool.AskAlways("Database username", controller.validateDatabaseUsername)
		}

		databasePassword = cmd_tool.Ask("Database password")
		if err := controller.validateDatabasePassword(databasePassword); err != nil {
			printer.Error(err)
			databasePassword = cmd_tool.AskAlways("Database password", controller.validateDatabasePassword)
		}

		databaseName = cmd_tool.Ask("Database name")
		if err := controller.validateDatabaseName(databaseName); err != nil {
			printer.Error(err)
			databaseName = cmd_tool.AskAlways("Database name", controller.validateDatabaseName)
		}
	}

	printer.Info("Create project: " + projectName + "...")

	// create not exist folders & files
	config := &models.ProjectConfig{
		Name:        projectName,
		Description: projectDescription,
		UseDatabase: useDatabase,

		Database: &models.DatabaseConfig{
			Server:   databaseServer,
			Port:     databasePort,
			Username: databaseUsername,
			Password: databasePassword,
			Name:     databaseName,
		},
	}
	if err := controller.skeleton.CreateObjects(config); err != nil {
		return err
	}

	// format project
	//if err := project_formatter.All("."); err != nil {
	//	printer.Error(err)
	//}

	return nil
}
