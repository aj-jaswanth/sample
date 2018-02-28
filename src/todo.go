package main

import (
	"github.com/aj-jaswanth/sample/src/database"
	"github.com/aj-jaswanth/sample/src/server"
)

func main() {
	database.RunMigrations()
	server.Serve()
}
