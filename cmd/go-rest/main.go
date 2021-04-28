package main

import (
	"log"

	"github.com/marcosap/go-rest/internal/animals"
	"github.com/marcosap/go-rest/internal/api"
	"github.com/marcosap/go-rest/internal/cars"
	"github.com/marcosap/go-rest/internal/database"
)

func main() {
	log.Println("I'm go-rest, nice to meet you!")

	db := database.NewDatabase()
	defer db.Disconnect()

	api := api.NewAPI()

	animals := animals.NewAnimalsResource(db)
	api.AddResource(animals)

	cars := cars.CarsResource{}
	api.AddResource(&cars)

	api.Start()
}
