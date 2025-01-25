package main

import (
	"myapp/handlers"

	quokka "github.com/ali-ahadi1105/Quokka"
)

type application struct {
	App      *quokka.Quokka
	Handlers *handlers.Handlers
}

func main() {
	q := initialApp()
	q.App.ListenAndServe()
}
