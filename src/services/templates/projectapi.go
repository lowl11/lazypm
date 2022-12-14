package templates

const (
	apiRoutes = `package api

import (
	"github.com/labstack/echo/v4"
	"<% project_name %>/src/controllers"
	"<% project_name %>/src/definition"
	"<% project_name %>/src/events"
	"<% project_name %>/src/repositories"
	"net/http"
)

func setRoutes(server *echo.Echo) {
	logger := definition.Logger

	// ивенты
	apiEvents, err := events.Get()
	if err != nil {
		logger.Fatal(err, "Initializing events error")
	}

	// репозитории
	apiRepositories, err := repositories.Get(apiEvents)
	if err != nil {
		logger.Fatal(err, "Connecting to database error")
	}

	// контроллеры
	apiControllers := controllers.Get(apiRepositories, apiEvents)

	server.GET("/health", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})
	server.RouteNotFound("*", apiControllers.Static.RouteNotFound)
}`

	apiMiddlewares = `package api

import "github.com/labstack/echo/v4"

func setMiddlewares(server *echo.Echo) {
	//
}

func setMiddlewaresToGroup(group *echo.Group) {
	//
}`

	apiServer = `package api

import (
	"<% project_name %>/src/definition"
)

func StartServer() {
	server := definition.Server

	// проставлять роуты
	setRoutes(server)

	// проставлять миддлвейры
	setMiddlewares(server)

	// запуск сервера
	server.Logger.Fatal(server.Start(definition.Config.Server.Port))
}`
)
