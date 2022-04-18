package main

import (
	"log"

	"github.com/alexgo92/GO_level_2/homework_7/task_1/internal/parse"
	"github.com/google/uuid"
)

type In struct {
	Name  string
	Code  uuid.UUID
	Price int64
}

func main() {
	var values = map[string]interface{}{
		"name":  "Mario",
		"code":  "b0d4ce5d-2757-4699-948c-cfa72ba94f86",
		"price": "55",
	}

	// Prod := &In{}
	Prod := 0
	err := parse.ParseValuesIn(Prod, values)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(Prod)
	}

}
