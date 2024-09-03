package main

import (
	"genshin/internal/app"
	_ "github.com/go-sql-driver/mysql"

)

func main() {
	app := app.NewApplication()

	app.StartServer()
}
