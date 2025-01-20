package main

import (
	"log"
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

	quo.InfoLog.Println("Debug value is set to: ", quo.Debug)
	quo.AppName = "firtsApp"

	app := &application{
		App: quo,
	}

	return app
}
