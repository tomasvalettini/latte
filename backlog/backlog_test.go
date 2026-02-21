package backlog_test

import (
	"os"
	"testing"

	"github.com/tomasvalettini/latte/assert"
	"github.com/tomasvalettini/latte/backlog"
	testutils "github.com/tomasvalettini/latte/test-utils"
)

const tmp = "tmp/"

func TestBacklogLogic(t *testing.T) {
	bl := backlog.NewBacklog(tmp + "latte/test.json")
	items := bl.Load()
	assert.Assert(len(items) == 0, "There should not be any items yet!")

	t.Cleanup(func() {
		os.RemoveAll(tmp)
	})

	// create items here :)
	item := backlog.Item{
		Id:   0,
		Text: "test item",
	}

	items = append(items, item)
	assert.Assert(len(items) != 0, "There should be items in the list!")

	bl.Save(items)
	items = bl.Load()
	assert.Assert(len(items) != 0, "There should be items in the list!")
}

func TestBacklogLogicFailingFile(t *testing.T) {
	testutils.RequireExit(t, "TestBacklogLogicFailingFile", testingFailingFile)
}

func testingFailingFile() {
	bl := backlog.NewBacklog(tmp)
	bl.Load()
}
