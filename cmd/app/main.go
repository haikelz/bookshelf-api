package main

import (
	"bookshelf/internal/server/routes"
	"bookshelf/internal/utils"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	utils.LoadEnv()
	routes.BookRoute()
}
