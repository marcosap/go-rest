package animals

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcosap/go-rest/internal/api"
)

type AnimalsResource struct{}

func (animalsResource *AnimalsResource) GetRoutes() []api.ApiRoute {

	return []api.ApiRoute{
		{
			Url:     "/api/animals",
			Handler: animalsResource.getAllAnimals,
			Method:  "GET",
		},
		{
			Url:     "/api/animals/{name}",
			Handler: animalsResource.getAnimalByName,
			Method:  "GET",
		},
	}
}

var MockAnimals = []Animal{
	{Name: "Bob", Type: "dog"},
	{Name: "Felix", Type: "cat"},
}

func (animalsResource *AnimalsResource) getAllAnimals(resp http.ResponseWriter, req *http.Request) {

	data, err := json.Marshal(MockAnimals)

	if err != nil {
		log.Printf("animalResource.getAllAnimals - Error: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.Write(data)
}

func (animalsResource *AnimalsResource) getAnimalByName(resp http.ResponseWriter, req *http.Request) {

	paramName := mux.Vars(req)["name"]

	for _, animal := range MockAnimals {

		if animal.Name == paramName {

			data, err := json.Marshal(animal)

			if err != nil {
				log.Printf("animalResource.getAllAnimals - Error: %s", err)
				resp.WriteHeader(http.StatusInternalServerError)
				return
			}

			resp.Write(data)
		}
	}

	resp.WriteHeader(http.StatusNotFound)
}
