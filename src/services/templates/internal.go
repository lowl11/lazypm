package templates

var (
	templates = map[string]string{
		// root
		"root_gitignore": rootGitignore,
		"root_readme":    rootReadme,
		"root_go_mod":    rootGoMod,
		"root_main":      rootMain,

		// config
		"config_debug":   configDebug,
		"config_release": configRelease,

		// controllers
		"controllers_base": controllersBase,

		// data
		// models
		"models_rest":  modelsRest,
		"models_error": modelsError,
	}
)
