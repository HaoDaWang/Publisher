package routes

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"publisher/utils"
)

type Card struct {
	Name string `json:"name"`
	Meta string `json:"meta"`
	Color string `json:"-"`
	Value int `json:"value"`
	Description string  `json:"description"`
}

type Envs struct {
	Cards []Card  `json:"envs"`
}

func Index(res http.ResponseWriter, req *http.Request){
	var envs Envs

	file, err := os.Open("config/env.json")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&envs)

	if err != nil {
		log.Fatal(err)
	}

	utils.SendTpl(
		res,
		[]string{"index", "cards", "terminals", "branch"},
		template.FuncMap{},
		envs,
	)
}