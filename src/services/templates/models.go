package templates

const (
	modelsRest = `package models

type Response struct {
	Status       string      ` + "`json:\"status\"`" + `
	InnerMessage string      ` + "`json:\"inner_message\"`" + `
	Message      string      ` + "`json:\"message\"`" + `
	Body         interface{} ` + "`json:\"body,omitempty\"`" + `
}`

	modelsError = `package models

type Error struct {
	TechMessage     string ` + "`json:\"tech_message\"`" + `
	BusinessMessage string ` + "`json:\"business_message\"`" + `
}

func (e *Error) Error() string {
	return e.TechMessage + " -> " + e.BusinessMessage
}

func (e *Error) With(err error) *Error {
	with := e.copy()
	with.TechMessage = err.Error() + " | " + with.TechMessage
	return with
}

func (e *Error) copy() *Error {
	errorCopy := &Error{
		TechMessage:     e.TechMessage,
		BusinessMessage: e.BusinessMessage,
	}
	return errorCopy
}`
)
