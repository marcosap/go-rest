package main

import (
	"log"

	"github.com/marcosap/go-rest/internal/api"
)

func main() {
	log.Println("I'm go-rest, nice to meet you!")

	api := api.NewAPI()
	api.Start()
}
