package main

import "bookshelf-api/internal/server"

func main() {
	server := server.New()

	server.RegisterFiberRoutes()
	server.Listen(":5000")
}
