package cars

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/marcosap/go-rest/internal/api"
)

type CarsResource struct{}

func (carsResource *CarsResource) GetRoutes() []api.ApiRoute {

	return []api.ApiRoute{
		{
			Url:     "/api/cars",
			Handler: carsResource.getAllCars,
			Method:  "GET",
		},
	}
}

var MockCars = []Car{
	{Model: "Tucker Torpedo", Year: "1948"},
	{Model: "Ford GT40", Year: "1966"},
}

func (carsResource *CarsResource) getAllCars(resp http.ResponseWriter, req *http.Request) {

	data, err := json.Marshal(MockCars)

	if err != nil {
		log.Printf("carsResource.getAllCars - Error: %s", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.Write(data)
}
