package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
}

func NewAPI() *API {
	return &API{
		router: mux.NewRouter(),
	}
}

func (a *API) Start() {

	log.Printf("Starting the API")

	err := http.ListenAndServe(":9000", a.router)

	if err != nil {
		log.Printf("API error: %s", err)
	}
}
