package templates

const (
	definitionConfig = `package definition

import (
	"github.com/lowl11/lazyconfig/confapi"
	"github.com/lowl11/lazylog/logapi"
	"os"
)

type Configuration struct {
	Server struct {
		Port string ` + "`json:\"port\"`" + `
	} ` + "`json:\"server\"`" + `

	<% definition_config_database %>
}

var Config Configuration
var Logger logapi.ILogger

func Init() {
	Config = Configuration{}
	isProduction := os.Getenv("env") == "production"

	logger := logapi.New().File("info", "logs")

	if err := confapi.Read(&Config, isProduction); err != nil {
		logger.Fatal(err, "Reading config error")
	}

	Logger = logger
	initServer()
}`

	definitionConstants = `package definition

const (
	//
)`

	definitionServer = `package definition

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"<% project_name %>/src/middlewares"
)

var Server *echo.Echo

func initServer() {
	Server = echo.New()

	Server.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	Server.Use(middleware.Secure())
	Server.Use(middleware.Recover())
	Server.Use(middlewares.Timeout())
}`

	definitionConfigDatabase = `Database struct {
		Connection     string ` + "`json:\"connection\"`" + `
		MaxConnections int    ` + "`json:\"max_connections\"`" + `
		Lifetime       int    ` + "`json:\"lifetime\"`" + `
	} ` + "`json:\"database\"`" + ``
)
