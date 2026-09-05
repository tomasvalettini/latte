package datamodel

import (
	"testing"

	"github.com/tomasvalettini/latte/assert"
	testutils "github.com/tomasvalettini/latte/test-utils"
)

func TestDripsGetUnusedId(t *testing.T) {
	drips := []Drip{}
	id := GetNextId(drips)
	assert.Assert(id == 0, "No drips, next id should be 0")

	drips = append(drips, Drip{Id: 1, Text: ""})
	id = GetNextId(drips)
	assert.Assert(id == 2, "Next id should be 2")

	drips = append(drips, Drip{Id: 3, Text: ""})
	id = GetNextId(drips)
	assert.Assert(id == 4, "Next id should be 4")
}

func TestMaxIdWidth(t *testing.T) {
	drips := []Drip{}
	w := MaxIdWidth(drips)
	assert.Assert(w == 0, "No drips, width should be 0!!!!!")

	drips = append(drips, Drip{Id: 1, Text: ""})
	w = MaxIdWidth(drips)
	assert.Assert(w == 1, "Width should be 1")

	drips = append(drips, Drip{Id: 11, Text: ""})
	w = MaxIdWidth(drips)
	assert.Assert(w == 2, "Width should be 2")

	drips = append(drips, Drip{Id: 111, Text: ""})
	w = MaxIdWidth(drips)
	assert.Assert(w == 3, "Width should be 3")

	drips = append(drips, Drip{Id: 1111, Text: ""})
	w = MaxIdWidth(drips)
	assert.Assert(w == 4, "Width should be 4")
}

func TestFindIndexFromId(t *testing.T) {
	drips := []Drip{}
	drips = append(drips, Drip{Id: 1, Text: ""})
	drips = append(drips, Drip{Id: 11, Text: ""})
	drips = append(drips, Drip{Id: 111, Text: ""})
	drips = append(drips, Drip{Id: 1111, Text: ""})

	id := FindIndexFromId(drips, 1)
	assert.Assert(id == 0, "Id in the wrong index")

	id = FindIndexFromId(drips, 11)
	assert.Assert(id == 1, "Id in the wrong index")

	id = FindIndexFromId(drips, 111)
	assert.Assert(id == 2, "Id in the wrong index")

	id = FindIndexFromId(drips, 1111)
	assert.Assert(id == 3, "Id in the wrong index")
}

func TestFindIndexFromIdNotFound(t *testing.T) {
	testutils.RequireExit(t, "TestFindIndexFromIdNotFound", testFindIndexFromIdFailing)
}

func testFindIndexFromIdFailing() {
	drips := []Drip{}
	drips = append(drips, Drip{Id: 1, Text: ""})
	drips = append(drips, Drip{Id: 11, Text: ""})

	FindIndexFromId(drips, 4)
}
