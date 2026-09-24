package datamodel

import (
	"fmt"
	"testing"

	"github.com/tomasvalettini/latte/assert"
)

func TestGetNextBlendId(t *testing.T) {
	blends := getTestBlends()
	nextId := GetNextBlendId(blends)

	assert.Assert(nextId == 5, "Blend id should be 5.")
}

func TestMaxBlendIdWidth(t *testing.T) {
	type testMaxBlendIdWidth struct {
		blends   []Blend
		expected int
	}

	testCases := []testMaxBlendIdWidth{
		{blends: []Blend{}, expected: 0},
		{blends: []Blend{{Id: 0, Title: "House Blend"}}, expected: 1},
		{blends: []Blend{{Id: 5, Title: "Blend 5"}}, expected: 1},
		{blends: []Blend{{Id: 10, Title: "Blend 10"}}, expected: 2},
		{blends: []Blend{{Id: 123, Title: "Blend 123"}}, expected: 3},
		{blends: []Blend{{Id: -1, Title: "Blend -1"}}, expected: 2},
		{blends: []Blend{{Id: -100, Title: "Blend -100"}}, expected: 4},
		{blends: []Blend{{Id: 9, Title: "Blend 9"}, {Id: 5, Title: "Blend 5"}}, expected: 1},
		{blends: []Blend{{Id: 100, Title: "Blend 100"}, {Id: 5, Title: "Blend 5"}}, expected: 3},
		{blends: []Blend{{Id: 7, Title: "Blend 7"}, {Id: 7, Title: "Blend 7b"}}, expected: 1},
		{blends: []Blend{{Id: -5, Title: "Blend -5"}, {Id: -3, Title: "Blend -3"}, {Id: -10, Title: "Blend -10"}}, expected: 3},
		{blends: []Blend{{Id: -5, Title: "Blend -5"}, {Id: 0, Title: "House Blend"}, {Id: 3, Title: "Blend 3"}}, expected: 2},
		{blends: getTestBlends(), expected: 1},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("blends=%v", tc.blends)
		t.Run(name, func(t *testing.T) {
			result := MaxBlendIdWidth(tc.blends)
			msg := fmt.Sprintf("%s,result=%d,expected=%d", name, result, tc.expected)
			assert.Assert(result == tc.expected, msg)
		})
	}
}

func TestSortBlendsById(t *testing.T) {
	type testSortBlendsById struct {
		name     string
		input    []Blend
		expected []int // expected order of Ids after sort
	}

	testCases := []testSortBlendsById{
		{
			name:     "empty",
			input:    []Blend{},
			expected: []int{},
		},
		{
			name:     "single",
			input:    []Blend{{Id: 5, Title: "Blend 5"}},
			expected: []int{5},
		},
		{
			name:     "already sorted",
			input:    []Blend{{Id: 1, Title: "A"}, {Id: 2, Title: "B"}, {Id: 3, Title: "C"}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "reverse sorted",
			input:    []Blend{{Id: 3, Title: "C"}, {Id: 2, Title: "B"}, {Id: 1, Title: "A"}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "duplicate ids",
			input:    []Blend{{Id: 2, Title: "B1"}, {Id: 1, Title: "A"}, {Id: 2, Title: "B2"}},
			expected: []int{1, 2, 2},
		},
		{
			name:     "negative ids",
			input:    []Blend{{Id: -5, Title: "Neg5"}, {Id: -1, Title: "Neg1"}, {Id: -10, Title: "Neg10"}},
			expected: []int{-10, -5, -1},
		},
		{
			name:     "mixed positive negative",
			input:    []Blend{{Id: 5, Title: "Pos5"}, {Id: -3, Title: "Neg3"}, {Id: 0, Title: "Zero"}},
			expected: []int{-3, 0, 5},
		},
		{
			name:     "from getTestBlends",
			input:    getTestBlends(),
			expected: []int{1, 2, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// make a copy to avoid mutating test case
			blends := make([]Blend, len(tc.input))
			copy(blends, tc.input)

			SortBlendsById(blends)

			if len(blends) != len(tc.expected) {
				assert.Assert(false, fmt.Sprintf("%s: length mismatch got %d want %d", tc.name, len(blends), len(tc.expected)))
				return
			}
			for i, id := range tc.expected {
				if blends[i].Id != id {
					assert.Assert(false, fmt.Sprintf("%s: at index %d got Id %d want %d", tc.name, i, blends[i].Id, id))
					return
				}
			}
		})
	}
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
