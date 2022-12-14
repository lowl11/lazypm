package main

import (
	"lazypm/src/api"
	"lazypm/src/definition"
)

func main() {
	definition.Init()
	api.RunCLI()
}
