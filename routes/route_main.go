package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"publisher/cli"
	"publisher/utils"
	"time"
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

func addTask(env int, branch string) bool {
	task := cli.Task{
		Name:fmt.Sprintf("部署%s", cli.TaskNameMap[env]),
		Branch: branch,
		Id: time.Now(),
		Env: env,
	}

	task.Start()

	return true
}

func removeTask(id time.Time){
	cli.Stop(id)
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
		[]string{"index", "cards", "terminals"},
		template.FuncMap{},
		envs,
	)
}