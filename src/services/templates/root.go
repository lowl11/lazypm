package templates

const (
	rootGoMod = `module <% project_name %>
	
go 1.18

require (
	<% gomod_sqlx %>
	<% gomod_postgres %>
	github.com/labstack/echo/v4 v4.9.1
	github.com/lowl11/lazyconfig v1.0.4
	github.com/lowl11/lazyfile v0.1.0
	github.com/lowl11/lazylog v1.1.1
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
)
`

	rootMain = `package main

import (
	"<% project_name %>/src/api"
	"<% project_name %>/src/definition"
)

func main() {
	definition.Init()
	api.StartServer()
}`

	rootGitignore = `.idea
*.exe
*.log
logs
main
app`

	rootReadme = `# <% project_name %>

> <% project_description %>`

	rootPackage = `{
	"project": {
		"name": "<% project_name %>",
		"is_database": <% project_is_database %>
	}
}`
)
