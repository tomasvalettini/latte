package dripdatasource

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/tomasvalettini/latte/assert"
	dripdatamodel "github.com/tomasvalettini/latte/drips/data/model"
)

type DripBacklog struct {
	dripsPath string
}

func NewDripBacklog(path string) *DripBacklog {
	return &DripBacklog{
		dripsPath: path,
	}
}

func (backlog *DripBacklog) Load() []dripdatamodel.Drip {
	data, err := os.ReadFile(backlog.dripsPath)

	if err != nil {
		if os.IsNotExist(err) {
			return []dripdatamodel.Drip{}
		}

		log.Fatalln("Error opening backlog file :(.")
	}

	var drips []dripdatamodel.Drip

	err = json.Unmarshal(data, &drips)
	assert.Assert(err == nil, "Error while parsing json.")

	return drips
}

func (backlog *DripBacklog) Save(drips []dripdatamodel.Drip) {
	err := os.MkdirAll(filepath.Dir(backlog.dripsPath), 0o755)
	assert.Assert(err == nil, "Error while creating and opening drip db.")

	data, err := json.MarshalIndent(drips, "", "  ")
	assert.Assert(err == nil, "Error while creating json.")

	data = append(data, '\n')
	os.WriteFile(backlog.dripsPath, data, 0o644)
}
