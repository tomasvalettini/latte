package taskdatasource

import (
	"os"
	"testing"

	"github.com/tomasvalettini/latte/assert"
	taskdatamodel "github.com/tomasvalettini/latte/tasks/data/model"
	taskpath "github.com/tomasvalettini/latte/tasks/path"
	testutils "github.com/tomasvalettini/latte/test-utils"
)

func TestTaskBacklogLogic(t *testing.T) {
	ttp := taskpath.GetTestingTaskPath()
	tbl := NewTaskBacklog(ttp.GetTaskPath())
	tasks := tbl.Load()
	assert.Assert(len(tasks) == 0, "There should not be any tasks yet!")

	t.Cleanup(func() {
		os.RemoveAll(taskpath.TMP)
	})

	// create tasks here :)
	task := taskdatamodel.Task{
		Id:   0,
		Text: "test task",
	}

	tasks = append(tasks, task)
	assert.Assert(len(tasks) != 0, "There should be tasks in the list!")

	tbl.Save(tasks)
	tasks = tbl.Load()
	assert.Assert(len(tasks) != 0, "There should be tasks in the list!")
}

func TestTaskBacklogLogicFailingFile(t *testing.T) {
	testutils.RequireExit(t, "TestTaskBacklogLogicFailingFile", testingFailingFile)
}

func testingFailingFile() {
	tbl := NewTaskBacklog(taskpath.TMP)
	tbl.Load()
}
