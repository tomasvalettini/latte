package backlog_test

import (
	"os"
	"testing"

	"github.com/tomasvalettini/latte/assert"
	"github.com/tomasvalettini/latte/backlog"
	testutils "github.com/tomasvalettini/latte/test-utils"
)

const tmp = "tmp/"

type TestTaskPath struct{}

func (ttp *TestTaskPath) GetTaskPath() string {
	return tmp + "latte/test.json"
}

func TestBacklogLogic(t *testing.T) {
	ttp := getTestingTaskPath()
	bl := backlog.NewBacklog(ttp.GetTaskPath())
	tasks := bl.Load()
	assert.Assert(len(tasks) == 0, "There should not be any tasks yet!")

	t.Cleanup(func() {
		os.RemoveAll(tmp)
	})

	// create tasks here :)
	task := backlog.Task{
		Id:   0,
		Text: "test task",
	}

	tasks = append(tasks, task)
	assert.Assert(len(tasks) != 0, "There should be tasks in the list!")

	bl.Save(tasks)
	tasks = bl.Load()
	assert.Assert(len(tasks) != 0, "There should be tasks in the list!")
}

func TestBacklogLogicFailingFile(t *testing.T) {
	testutils.RequireExit(t, "TestBacklogLogicFailingFile", testingFailingFile)
}

func testingFailingFile() {
	bl := backlog.NewBacklog(tmp)
	bl.Load()
}

func getTestingTaskPath() backlog.TaskPath {
	return &TestTaskPath{}
}
