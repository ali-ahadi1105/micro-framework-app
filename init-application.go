package main

import (
	"log"
	"myapp/handlers"
	"os"

	quokka "github.com/ali-ahadi1105/Quokka"
)

func initialApp() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	quo := &quokka.Quokka{}
	err = quo.New(path)
	if err != nil {
		log.Fatal(err)
	}

	quo.AppName = "firtsApp"

	myHandlers := &handlers.Handlers{
		App: quo,
	}

	app := &application{
		App:      quo,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
