package backlog

import (
	"testing"

	"github.com/tomasvalettini/latte/assert"
	testutils "github.com/tomasvalettini/latte/test-utils"
)

func TestTasksGetUnusedId(t *testing.T) {
	tasks := []Task{}
	id := GetNextId(tasks)
	assert.Assert(id == 0, "No tasks, next id should be 0")

	tasks = append(tasks, Task{Id: 1, Text: ""})
	id = GetNextId(tasks)
	assert.Assert(id == 2, "Next id should be 2")

	tasks = append(tasks, Task{Id: 3, Text: ""})
	id = GetNextId(tasks)
	assert.Assert(id == 4, "Next id should be 4")
}

func TestMaxIdWidth(t *testing.T) {
	tasks := []Task{}
	w := MaxIdWidth(tasks)
	assert.Assert(w == 0, "No tasks, width should be 0!!!!!")

	tasks = append(tasks, Task{Id: 1, Text: ""})
	w = MaxIdWidth(tasks)
	assert.Assert(w == 1, "Width should be 1")

	tasks = append(tasks, Task{Id: 11, Text: ""})
	w = MaxIdWidth(tasks)
	assert.Assert(w == 2, "Width should be 2")

	tasks = append(tasks, Task{Id: 111, Text: ""})
	w = MaxIdWidth(tasks)
	assert.Assert(w == 3, "Width should be 3")

	tasks = append(tasks, Task{Id: 1111, Text: ""})
	w = MaxIdWidth(tasks)
	assert.Assert(w == 4, "Width should be 4")
}

func TestFindIndexFromId(t *testing.T) {
	tasks := []Task{}
	tasks = append(tasks, Task{Id: 1, Text: ""})
	tasks = append(tasks, Task{Id: 11, Text: ""})
	tasks = append(tasks, Task{Id: 111, Text: ""})
	tasks = append(tasks, Task{Id: 1111, Text: ""})

	id := FindIndexFromId(tasks, 1)
	assert.Assert(id == 0, "Id in the wrong index")

	id = FindIndexFromId(tasks, 11)
	assert.Assert(id == 1, "Id in the wrong index")

	id = FindIndexFromId(tasks, 111)
	assert.Assert(id == 2, "Id in the wrong index")

	id = FindIndexFromId(tasks, 1111)
	assert.Assert(id == 3, "Id in the wrong index")
}

func TestFindIndexFromIdNotFound(t *testing.T) {
	testutils.RequireExit(t, "TestFindIndexFromIdNotFound", testFindIndexFromIdFailing)
}

func testFindIndexFromIdFailing() {
	tasks := []Task{}
	tasks = append(tasks, Task{Id: 1, Text: ""})
	tasks = append(tasks, Task{Id: 11, Text: ""})

	FindIndexFromId(tasks, 4)
}
