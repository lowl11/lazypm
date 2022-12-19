package templates

var (
	templates = map[string]string{
		// root
		"root_gitignore": rootGitignore,
		"root_readme":    rootReadme,
		"root_go_mod":    rootGoMod,
		"root_main":      rootMain,
		"root_package":   rootPackage,

		// config
		"config_debug":    configDebug,
		"config_release":  configRelease,
		"config_database": configDatabase,

		// api
		"api_routes":              apiRoutes,
		"api_middlewares":         apiMiddlewares,
		"api_server":              apiServer,
		"api_import_repositories": apiImportRepositories,
		"api_repositories_create": apiRepositoriesCreate,
		"api_field_repositories":  apiFieldRepositories,

		// definition
		"definition_config":          definitionConfig,
		"definition_constants":       definitionConstants,
		"definition_server":          definitionServer,
		"definition_config_database": definitionConfigDatabase,

		// middlewares
		"middlewares_timeout": middlewaresTimeout,

		// controllers
		"controllers_api":  controllersApi,
		"controllers_base": controllersBase,
		// static controller
		"controllers_static_api":        controllersStaticApi,
		"controllers_static_controller": controllersStaticController,
		"controllers_static_validation": controllersStaticValidation,
		// new controller
		"controllers_new_api":        controllersNewApi,
		"controllers_new_controller": controllersNewController,
		"controllers_new_validation": controllersNewValidation,

		// repositories
		"repositories_api": repositoriesApi,
		"repository_base":  repositoryBase,

		// data
		// models
		"models_rest":  modelsRest,
		"models_error": modelsError,

		// errors
		"errors_static": errorsStatic,

		// events
		"events_api": eventsApi,

		// script
		"events_script_api":                 eventsScriptApi,
		"events_script_event":               eventsScriptEvent,
		"events_script_internal":            eventsScriptInternal,
		"events_script_create":              eventsScriptCreate,
		"events_script_contain":             eventsScriptContain,
		"controllers_import_repositories":   controllersImportRepositories,
		"controllers_argument_repositories": controllersArgumentRepositories,
	}
)
