package main

import (
	"todo-api/pkg/app"
	"todo-api/pkg/config"
)

func main() {
	// load config
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// run app
	app := app.NewApp(cfg)
	app.Run()
}
