package animals

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcosap/go-rest/internal/api"
	"github.com/marcosap/go-rest/internal/database"
)

type AnimalsResource struct {
	db *database.Database
}

func NewAnimalsResource(db *database.Database) *AnimalsResource {
	return &AnimalsResource{
		db: db,
	}
}

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
		{
			Url:     "/api/animals",
			Handler: animalsResource.createAnimal,
			Method:  "POST",
		},
		{
			Url:     "/api/animals/{name}",
			Handler: animalsResource.updateAnimalByName,
			Method:  "PUT",
		},
		{
			Url:     "/api/animals/{name}",
			Handler: animalsResource.deleteAnimalByName,
			Method:  "DELETE",
		},
	}
}

var MockAnimals = []Animal{
	{Name: "Bob", Type: "dog"},
	{Name: "Felix", Type: "cat"},
}

func (animalsResource *AnimalsResource) getAllAnimals(resp http.ResponseWriter, req *http.Request) {

	animals, err := animalsResource.db.RetrieveAll(&Animal{})

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(animals)

	if err != nil {
		log.Printf("animalResource.getAllAnimals - Error: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.Write(data)
}

func (animalsResource *AnimalsResource) getAnimalByName(resp http.ResponseWriter, req *http.Request) {

	paramName := mux.Vars(req)["name"]

	animal := Animal{Name: paramName}

	err := animalsResource.db.RetrieveOne(&animal)

	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(animal)

	if err != nil {
		log.Printf("animalResource.getAnimalByName - Error: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.Write(data)

}

func (animalsResource *AnimalsResource) createAnimal(resp http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	var animal Animal

	json.Unmarshal(body, &animal)

	err = animalsResource.db.Create(&animal)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (animalsResource *AnimalsResource) updateAnimalByName(resp http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	var animal Animal
	err = json.Unmarshal(body, &animal)

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	paramName := mux.Vars(req)["name"]

	animal.Name = paramName

	exists, err := animalsResource.db.Update(&animal)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !exists {
		resp.WriteHeader(http.StatusNotFound)
	}

}

func (animalsResource *AnimalsResource) deleteAnimalByName(resp http.ResponseWriter, req *http.Request) {

	paramName := mux.Vars(req)["name"]

	animal := Animal{Name: paramName}

	exists, err := animalsResource.db.Delete(&animal)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !exists {
		resp.WriteHeader(http.StatusNotFound)
	}

}
