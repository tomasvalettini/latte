package backlog_test

import (
	"os"
	"testing"

	"github.com/tomasvalettini/latte/assert"
	"github.com/tomasvalettini/latte/backlog"
	testutils "github.com/tomasvalettini/latte/test-utils"
)

func TestBacklogLogic(t *testing.T) {
	ttp := backlog.GetTestingTaskPath()
	bl := backlog.NewBacklog(ttp.GetTaskPath())
	tasks := bl.Load()
	assert.Assert(len(tasks) == 0, "There should not be any tasks yet!")

	t.Cleanup(func() {
		os.RemoveAll(backlog.TMP)
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
	bl := backlog.NewBacklog(backlog.TMP)
	bl.Load()
}
