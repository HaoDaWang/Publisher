package cli

import (
	"fmt"
)

var tasks []Task
var TaskNameMap = map[int]string {
	TEST_ENV_1: "测试环境1",
	TEST_ENV_2: "测试环境2",
	LOCAL_TEST_ENV: "本地环境",
}

type Task struct {
	Id int64
	Name string
	Env int
	Branch string
	Log []string
}

func (task *Task) Start() int64 {
	task.Name = fmt.Sprintf("部署%s-%s", TaskNameMap[task.Env], task.Branch)
	task.Log = Publish(task.Branch, task.Env)

	tasks = append(tasks, *task)

	return task.Id
}

func Stop(_id int64) {
	index := -1

	for i, task := range tasks  {
		if task.Id == _id {
			index = i
		}
	}

	if index == -1 {
		return
	}

	tasks = append(tasks[:index], tasks[index + 1:]...)
}

func GetTasks() []Task {
	return tasks
}