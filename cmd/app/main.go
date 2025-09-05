package main

import (
	"bookshelf/internal/configs"
	"bookshelf/internal/server/routes"
	"bookshelf/internal/utils"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	configs.NewViper()
	utils.LoadEnv()

	echo := configs.NewEcho()
	db := configs.NewGorm()

	routes.BookRoute(db, echo)
}
