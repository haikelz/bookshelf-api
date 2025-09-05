package main

import (
	"bookshelf/internal/configs"
	"bookshelf/internal/server/routes"
	"bookshelf/internal/utils"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title Bookshelf API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	configs.NewViper()
	utils.LoadEnv()

	echo := configs.NewEcho()
	db := configs.NewGorm()

	routes.BookRoute(db, echo)
}
