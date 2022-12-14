package templates

const (
	repositoriesApi = `package repositories

import (
	"github.com/jmoiron/sqlx"
	"<% project_name %>/src/definition"
	"<% project_name %>/src/events"
	"time"

	_ "github.com/lib/pq"
)

type ApiRepositories struct {
	//
}

func Get(apiEvents *events.ApiEvents) (*ApiRepositories, error) {
	config := definition.Config.Database
	logger := definition.Logger

	connection, err := sqlx.Open("postgres", config.Connection)
	if err != nil {
		return nil, err
	}

	connection.SetMaxOpenConns(config.MaxConnections)
	connection.SetMaxIdleConns(config.MaxConnections)
	connection.SetConnMaxIdleTime(time.Duration(config.Lifetime) * time.Minute)

	logger.Info("Ping database...")
	if err = connection.Ping(); err != nil {
		return nil, err
	}
	logger.Info("Ping database done!")

	logger.Info("Initialization database...")
	defer logger.Info("Initialization database done!")
	return &ApiRepositories{
		//
	}, nil
}`

	repositoryBase = `package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"<% project_name %>/src/definition"
	"<% project_name %>/src/events/script_event"
	"strings"
	"time"
)

const (
	requiredFieldError = "Field '%s' is required"
)

type Base struct {
	script *script_event.Event
}

func CreateBase(script *script_event.Event) Base {
	return Base{
		script: script,
	}
}

func (repo *Base) StartScript(name string) string {
	return repo.script.StartScript(name)
}

func (repo *Base) Script(folder, name string) string {
	return repo.script.Script(folder, name)
}

func (repo *Base) Ctx(customTimeout ...time.Duration) (context.Context, func()) {
	defaultTimeout := time.Second * 5
	if len(customTimeout) > 0 {
		defaultTimeout = customTimeout[0]
	}
	return context.WithTimeout(context.Background(), defaultTimeout)
}

func (repo *Base) CloseRows(rows *sqlx.Rows) {
	if err := rows.Close(); err != nil {
		definition.Server.Logger.Error("Closing rows error", err)
	}
}

func (repo *Base) Rollback(transaction *sqlx.Tx) {
	if err := transaction.Rollback(); err != nil {
		if !strings.Contains(err.Error(), "sql: transaction has already been committed or rolled back") {
			definition.Logger.Error(err, "Rollback transaction error")
		}
	}
}

func (repo *Base) Transaction(connection *sqlx.DB, transactionActions func(tx *sqlx.Tx) error) error {
	transaction, err := connection.Beginx()
	if err != nil {
		return err
	}
	defer repo.Rollback(transaction)

	if err = transactionActions(transaction); err != nil {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo *Base) RequiredField(value interface{}, name string) error {
	if value == nil {
		return errors.New(fmt.Sprintf(requiredFieldError, name))
	}

	_, isString := value.(string)
	if isString {
		if value.(string) == "" {
			return errors.New(fmt.Sprintf(requiredFieldError, name))
		}
	}

	_, isInt := value.(int)
	if isInt {
		if value.(int) == 0 {
			return errors.New(fmt.Sprintf(requiredFieldError, name))
		}
	}
	return nil
}`
)
