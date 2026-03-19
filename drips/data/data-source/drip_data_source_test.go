package dripdatasource

import (
	"os"
	"testing"

	"github.com/tomasvalettini/latte/assert"
	dripdatamodel "github.com/tomasvalettini/latte/drips/data/model"
	drippath "github.com/tomasvalettini/latte/drips/path"
	testutils "github.com/tomasvalettini/latte/test-utils"
)

func TestDripBacklogLogic(t *testing.T) {
	ttp := drippath.GetTestingDripPath()
	tbl := NewDripBacklog(ttp.GetDripPath())
	drips := tbl.Load()
	assert.Assert(len(drips) == 0, "There should not be any drips yet!")

	t.Cleanup(func() {
		os.RemoveAll(drippath.TMP)
	})

	// create drips here :)
	drip := dripdatamodel.Drip{
		Id:   0,
		Text: "test drip",
	}

	drips = append(drips, drip)
	assert.Assert(len(drips) != 0, "There should be drips in the list!")

	tbl.Save(drips)
	drips = tbl.Load()
	assert.Assert(len(drips) != 0, "There should be drips in the list!")
}

func TestDripBacklogLogicFailingFile(t *testing.T) {
	testutils.RequireExit(t, "TestDripBacklogLogicFailingFile", testingFailingFile)
}

func testingFailingFile() {
	tbl := NewDripBacklog(drippath.TMP)
	tbl.Load()
}
