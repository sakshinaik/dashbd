package main

import (
	"dashboard-api/config"
)

func main() {
	config.Setup()

	app := App{}
	app.InitApp()
	defer app.DB.Close()
	app.Run()
}
