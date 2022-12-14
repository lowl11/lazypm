package api

import (
	"github.com/lowl11/lazy-cli/cliapi"
	"lazypm/src/definition"
)

func RunCLI() {
	// create CLI client
	cli := cliapi.New()

	// routes
	setRoutes(cli)

	// run CLI
	if err := cli.Start(); err != nil {
		definition.Printer.Error(err)
	}
}
