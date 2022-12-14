package api

import (
	"github.com/lowl11/lazy-cli/cliapi"
	"lazypm/src/controllers"
	"lazypm/src/controllers/controller_controller"
	"lazypm/src/controllers/project_controller"
	"lazypm/src/definition"
	"lazypm/src/events"
)

func setRoutes(cli *cliapi.Cli) {
	basePath := "../test-service"
	//basePath := ""

	apiEvents, err := events.Get(basePath)
	if err != nil {
		definition.Printer.Fatal(err)
	}

	apiControllers := controllers.Get(apiEvents)

	setProject(cli, apiControllers.Project)
	setController(cli, apiControllers.Controller)
}

func setProject(cli *cliapi.Cli, controller *project_controller.Controller) {
	cli.Route(controller.Init, "project", "init")
}

func setController(cli *cliapi.Cli, controller *controller_controller.Controller) {
	cli.Route(controller.New, "controller", "new")
}
