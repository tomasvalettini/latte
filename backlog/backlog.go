package backlog

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/tomasvalettini/latte/assert"
)

type Backlog struct {
	tasksPath string
}

func NewBacklog(path string) *Backlog {
	return &Backlog{
		tasksPath: path,
	}
}

func (backlog *Backlog) Load() []Task {
	data, err := os.ReadFile(backlog.tasksPath)

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}
		}

		log.Fatalln("Error opening backlog file :(.")
	}

	var tasks []Task

	err = json.Unmarshal(data, &tasks)
	assert.Assert(err == nil, "Error while parsing json.")

	return tasks
}

func (backlog *Backlog) Save(tasks []Task) {
	err := os.MkdirAll(filepath.Dir(backlog.tasksPath), 0o755)
	assert.Assert(err == nil, "Error while creating and opening task db.")

	data, err := json.MarshalIndent(tasks, "", "  ")
	assert.Assert(err == nil, "Error while creating json.")

	data = append(data, '\n')
	os.WriteFile(backlog.tasksPath, data, 0o644)
}
