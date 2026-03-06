package controller

import (
	"os"
	"testing"

	taskdatasource "github.com/tomasvalettini/latte/tasks/data/data-source"
	taskpath "github.com/tomasvalettini/latte/tasks/path"
)

func GetTestTaskController() *TaskController {
	tp := taskpath.GetTestingTaskPath()
	bl := taskdatasource.NewTaskBacklog(tp.GetTaskPath())
	return &TaskController{
		taskPath:   tp,
		dataSource: bl,
	}
}

func TestTaskController(t *testing.T) {
	tc := GetTestTaskController()

	tc.ListTasks()

	tc.AddTask("test")
	tc.AddTask("test")
	tc.AddTask("test")
	tc.AddTask("test")
	tc.AddTask("test")

	tc.ListTasks()

	tc.DeleteTask("1")
	tc.DeleteTask("0")

	tc.UpdateTask("2", "new test")

	t.Cleanup(func() {
		os.RemoveAll(taskpath.TMP)
	})
}
