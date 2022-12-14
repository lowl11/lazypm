package cmd_tool

import (
	"bufio"
	"fmt"
	"lazypm/src/definition"
	"log"
	"os"
	"strings"
)

func Ask(question string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s: ", question)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		return strings.ToLower(strings.TrimSpace(response))
	}

	return ""
}

func AskAlways(question string, validationFunc func(value string) error) string {
	newValue := Ask(question)
	if err := validationFunc(newValue); err != nil {
		definition.Printer.Error(err)
		return AskAlways(question, validationFunc)
	}

	return newValue
}

func Confirm(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", question)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}

	return false
}
