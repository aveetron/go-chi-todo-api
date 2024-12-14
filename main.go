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
	app, err := app.NewApp(cfg)
	if err != nil {
		panic(err)
	}
	app.Run()
}
