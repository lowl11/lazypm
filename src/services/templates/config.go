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

	ConfigDatabase = `,
	"database": {
		"connection": "host=185.22.67.118 port=5432 user=testuser password=qwerty dbname=learnstest sslmode=disable",
    	"max_connections": 30,
    	"lifetime": 5	
	}`
)
