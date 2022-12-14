package templates

const (
	configDebug = `{
	"server": {
		"port": ":<% port %>"
	}<% config_database %>
}`
	configRelease = `{
	"server": {
		"port": ":<% port %>"
	}<% config_database %>
}`

	configDatabase = `,
	"database": {
		"connection": "host=<% database_server %> port=<% database_port %> user=<% database_user %> password=<% database_password %> dbname=<% database_name %> sslmode=disable",
    	"max_connections": 30,
    	"lifetime": 5	
	}`
)
