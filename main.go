package main

import (
	"ClassManagement/api"
	"ClassManagement/database"
)

/// @title API for Class Management
// @version 1.0
// @description These are APIs for Class, Subject and Teacher Management.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support (Ashwani)
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 0.0.0.0:8080
// @BasePath /
// @schemes http
func main() {
	database.ConnectDB()
	database.CreateTables()
	api.StartApi()
}
