package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"publisher/cli"
	"time"
)

type TasksPayload struct {
	Tasks []cli.Task
}

func AddTask(res http.ResponseWriter, req *http.Request){
	err := req.ParseForm()
	task := cli.Task{
		Id:time.Now(),
	}

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	err = decoder.Decode(task)

	if err != nil {
		log.Fatal(err)
	}

	task.Start()

	fmt.Fprintln(res, "done")
}

func RemoveTask(res http.ResponseWriter, req *http.Request){

}

func GetTasks(res http.ResponseWriter, req *http.Request){
	tasks := cli.GetTasks()

	resBody, err := json.Marshal(Response{
		payload: TasksPayload{
			Tasks:tasks,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks, resBody)

	fmt.Fprintln(res, tasks)
}