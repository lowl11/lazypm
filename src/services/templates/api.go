package templates

import "strings"

func Get(name string) string {
	return templates[name]
}

func FillVariables(name string, variables map[string]string) string {
	template := Get(name)
	for key, value := range variables {
		variableKey := "<% " + key + " %>"
		template = strings.ReplaceAll(template, variableKey, value)
	}
	return template
}
