package main

import (
	"log"

	"github.com/marcosap/go-rest/internal/animals"
	"github.com/marcosap/go-rest/internal/api"
	"github.com/marcosap/go-rest/internal/cars"
)

func main() {
	log.Println("I'm go-rest, nice to meet you!")

	api := api.NewAPI()

	animals := animals.AnimalsResource{}
	api.AddResource(&animals)

	cars := cars.CarsResource{}
	api.AddResource(&cars)

	api.Start()
}
