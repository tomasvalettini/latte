package backlog

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/tomasvalettini/latte/assert"
)

type Backlog struct {
	itemsPath string
}

func NewBacklog(path string) *Backlog {
	return &Backlog{
		itemsPath: path,
	}
}

func (backlog *Backlog) Load() []Item {
	data, err := os.ReadFile(backlog.itemsPath)

	if err != nil {
		if os.IsNotExist(err) {
			return []Item{}
		}

		log.Fatalln("Error opening backlog file :(.")
	}

	var items []Item

	err = json.Unmarshal(data, &items)
	assert.Assert(err == nil, "Error while parsing json.")

	return items
}

func (backlog *Backlog) Save(items []Item) {
	err := os.MkdirAll(filepath.Dir(backlog.itemsPath), 0o755)
	assert.Assert(err == nil, "Error while creating and opening task db.")

	data, err := json.MarshalIndent(items, "", "  ")
	assert.Assert(err == nil, "Error while creating json.")

	data = append(data, '\n')
	os.WriteFile(backlog.itemsPath, data, 0o644)
}
