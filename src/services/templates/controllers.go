package templates

const (
	// api
	controllersApi = `package controllers

import (
	"<% project_name %>/src/controllers/static_controller"
	"<% project_name %>/src/events"
	<% controllers_import_repositories %>
)

type ApiControllers struct {
	Static *static_controller.Controller
}

func Get(<% controllers_argument_repositories %>apiEvents *events.ApiEvents) *ApiControllers {
	return &ApiControllers{
		Static: static_controller.Create(),
	}
}`

	// base controller
	controllersBase = `package controller

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"<% project_name %>/src/data/models"
	"net/http"
)

type Base struct {
	//
}

func (controller *Base) Error(ctx echo.Context, err *models.Error, status ...int) error {
	responseStatus := http.StatusInternalServerError
	if len(status) > 0 {
		responseStatus = status[0]
	}

	errorObject := &models.Response{
		Status:       "ERROR",
		Message:      err.BusinessMessage,
		InnerMessage: err.TechMessage,
	}
	return ctx.JSON(responseStatus, errorObject)
}

func (controller *Base) NotFound(ctx echo.Context, err *models.Error) error {
	errorObject := &models.Response{
		Status:       "ERROR",
		Message:      err.BusinessMessage,
		InnerMessage: err.TechMessage,
	}
	return ctx.JSON(http.StatusNotFound, errorObject)
}

func (controller *Base) Ok(ctx echo.Context, response interface{}, messages ...string) error {
	defaultMessage := "OK"
	if len(messages) > 0 {
		defaultMessage = messages[0]
	}

	successObject := &models.Response{
		Status:       "OK",
		Message:      defaultMessage,
		InnerMessage: defaultMessage,
		Body:         response,
	}
	return ctx.JSON(http.StatusOK, successObject)
}

func (controller *Base) RequiredField(value interface{}, name string) error {
	if value == nil {
		return errors.New(fmt.Sprintf("Field %s is null, but it's required", name))
	}

	_, isString := value.(string)
	if isString && value.(string) == "" {
		return errors.New(fmt.Sprintf("Field %s is null or empty, but it's required", name))
	}

	_, isInt := value.(int)
	if isInt && value.(int) == 0 {
		return errors.New(fmt.Sprintf("Field %s is null or zero, but it's required", name))
	}

	return nil
}`

	// static controller
	controllersStaticApi = `package static_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"<% project_name %>/src/data/errors"
)

func (controller *Controller) Health(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

func (controller *Controller) RouteNotFound(ctx echo.Context) error {
	return controller.NotFound(ctx, errors.RouteNotFound)
}`
	controllersStaticController = `package static_controller

import "<% project_name %>/src/controllers/controller"

type Controller struct {
	controller.Base
}

func Create() *Controller {
	return &Controller{}
}`
	controllersStaticValidation = `package static_controller

//`

	controllersImportRepositories   = `"<% project_name %>/src/repositories"`
	controllersArgumentRepositories = "apiRepositories *repositories.ApiRepositories, "
)
