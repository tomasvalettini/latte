package taskdatasource

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/tomasvalettini/latte/assert"
	taskdatamodel "github.com/tomasvalettini/latte/tasks/data/model"
)

type TaskBacklog struct {
	tasksPath string
}

func NewTaskBacklog(path string) *TaskBacklog {
	return &TaskBacklog{
		tasksPath: path,
	}
}

func (backlog *TaskBacklog) Load() []taskdatamodel.Task {
	data, err := os.ReadFile(backlog.tasksPath)

	if err != nil {
		if os.IsNotExist(err) {
			return []taskdatamodel.Task{}
		}

		log.Fatalln("Error opening backlog file :(.")
	}

	var tasks []taskdatamodel.Task

	err = json.Unmarshal(data, &tasks)
	assert.Assert(err == nil, "Error while parsing json.")

	return tasks
}

func (backlog *TaskBacklog) Save(tasks []taskdatamodel.Task) {
	err := os.MkdirAll(filepath.Dir(backlog.tasksPath), 0o755)
	assert.Assert(err == nil, "Error while creating and opening task db.")

	data, err := json.MarshalIndent(tasks, "", "  ")
	assert.Assert(err == nil, "Error while creating json.")

	data = append(data, '\n')
	os.WriteFile(backlog.tasksPath, data, 0o644)
}
