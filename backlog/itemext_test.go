package backlog

import (
	"testing"

	"github.com/tomasvalettini/latte/assert"
)

func TestItemsGetUnusedId(t *testing.T) {
	items := []Item{}
	id := GetNextId(items)
	assert.Assert(id == 0, "No items, next id should be 0")

	items = append(items, Item{ Id: 1, Text: ""})
	id = GetNextId(items)
	assert.Assert(id == 2, "Next id should be 2")

	items = append(items, Item{ Id: 3, Text: ""})
	id = GetNextId(items)
	assert.Assert(id == 4, "Next id should be 3")
}

func TestMaxIdWidth(t *testing.T) {
	items := []Item{}
	w := MaxIdWidth(items)
	assert.Assert(w == 0, "No items, width should be 0!!!!!")

	items = append(items, Item{ Id: 1, Text: ""})
	w = MaxIdWidth(items)
	assert.Assert(w == 1, "Width should be 1")

	items = append(items, Item{ Id: 11, Text: ""})
	w = MaxIdWidth(items)
	assert.Assert(w == 2, "Width should be 2")

	items = append(items, Item{ Id: 111, Text: ""})
	w = MaxIdWidth(items)
	assert.Assert(w == 3, "Width should be 3")

	items = append(items, Item{ Id: 1111, Text: ""})
	w = MaxIdWidth(items)
	assert.Assert(w == 4, "Width should be 4")
}

