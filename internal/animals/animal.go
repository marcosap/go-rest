package animals

import "github.com/marcosap/go-rest/internal/database"

type Animal struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func (a *Animal) GetCollectionName() string {
	return "animals"
}

func (a *Animal) New() database.DatabaseEntity {
	return &Animal{}
}

func (a *Animal) GetFilterOne() map[string]interface{} {
	return map[string]interface{}{"name": a.Name}
}
