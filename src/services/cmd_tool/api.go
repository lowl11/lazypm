package cmd_tool

import (
	"bufio"
	"fmt"
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
