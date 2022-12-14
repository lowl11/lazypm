package templates

const (
	rootGoMod = `module <% project_name %>
	
go 1.18`

	rootMain = `package main

import "fmt"

func main() {
	fmt.Println("hello world")
}`

	rootGitignore = `.idea
*.exe
*.log
logs
main
app`

	rootReadme = `# <% project_name %>

> <% project_description %>`
)
