package definition

import (
	"fmt"
	"os"
)

type ConsolePrinter struct {
	//
}

func newPrinter() *ConsolePrinter {
	return &ConsolePrinter{}
}

func (printer *ConsolePrinter) Info(message string) {
	fmt.Println(message)
}

func (printer *ConsolePrinter) Error(err error) {
	if err == nil {
		return
	}

	fmt.Println(err.Error())
}

func (printer *ConsolePrinter) Fatal(err error) {
	if err == nil {
		return
	}

	fmt.Println(err.Error())
	os.Exit(1)
}
