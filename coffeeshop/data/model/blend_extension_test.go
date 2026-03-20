package datamodel

import (
	"testing"

	"github.com/tomasvalettini/latte/assert"
)

func TestGetNextBlendId(t *testing.T) {
	blends := getTestBlends()
	nextId := GetNextBlendId(blends)

	assert.Assert(nextId == 5, "Blend id should be 5.")
}

func getTestBlends() []Blend {
	return []Blend{
		{
			Id:    1,
			Title: "Blend 1",
			Drips: []Drip{
				{Id: 0, Text: "Lorem ipsum dolor sit amet."},
				{Id: 1, Text: "Lorem ipsum dolor sit amet."},
				{Id: 2, Text: "Lorem ipsum dolor sit amet."},
			},
		},
		{
			Id:    2,
			Title: "Blend 2",
			Drips: []Drip{
				{Id: 0, Text: "Lorem ipsum dolor sit amet."},
				{Id: 1, Text: "Lorem ipsum dolor sit amet."},
				{Id: 2, Text: "Lorem ipsum dolor sit amet."},
			},
		},
		{
			Id:    4,
			Title: "Empty Blend",
			Drips: []Drip{}, // Good for testing edge cases where a blend has no drips
		},
	}
}
