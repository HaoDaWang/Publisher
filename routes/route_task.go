package routes

import (
	"fmt"
	"log"
	"net/http"
	"publisher/JSON"
	"publisher/cli"
	"strconv"
	"time"
)

type TasksPayload struct {
	Tasks []cli.Task
}

func AddTask(res http.ResponseWriter, req *http.Request){
	err := req.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	branch := req.PostFormValue("branch")
	env := req.PostFormValue("env")
	envI, err := strconv.Atoi(env)

	if err != nil {
		log.Fatal(err)
	}

	task := cli.Task{
		Id: time.Now().Unix(),
		Env: envI,
		Branch: branch,
	}

	task.Start()

	fmt.Fprintln(res, "done")
}

func RemoveTask(res http.ResponseWriter, req *http.Request){
	err := req.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	id, err := strconv.ParseInt(req.PostFormValue("id"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	cli.Stop(id)

	fmt.Fprintln(res, "done")
}

func GetTasks(res http.ResponseWriter, req *http.Request){
	tasks := cli.GetTasks()

	json := JSON.H{
		"tasks": tasks,
	}

	fmt.Fprintln(res, json.Stringify())
}