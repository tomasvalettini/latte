package controller

import (
	"os"
	"testing"

	"github.com/tomasvalettini/latte/backlog"
)

func GetTestTaskController() *TaskController {
	tp := backlog.GetTestingTaskPath()
	bl := backlog.NewBacklog(tp.GetTaskPath())
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
		os.RemoveAll(backlog.TMP)
	})
}
