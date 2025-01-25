package handlers

import (
	"net/http"

	quokka "github.com/ali-ahadi1105/Quokka"
)

type Handlers struct {
	App *quokka.Quokka
}

func (handlers *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := handlers.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		handlers.App.ErrorLog.Println("Error in rendering: ", err)
	}
}
