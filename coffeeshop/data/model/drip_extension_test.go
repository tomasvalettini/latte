package datamodel

import (
	"fmt"

	"testing"

	"github.com/tomasvalettini/latte/assert"
)

func TestDripsGetUnusedId(t *testing.T) {
	type testGetNextId struct {
		drips    []Drip
		expected int
	}

	testCases := []testGetNextId{
		{drips: []Drip{}, expected: 0},
		{drips: []Drip{{Id: 1, Text: ""}}, expected: 2},
		{drips: []Drip{{Id: 1, Text: ""}, {Id: 3, Text: ""}}, expected: 4},
		{drips: []Drip{{Id: 0, Text: ""}, {Id: 2, Text: ""}}, expected: 3},
		{drips: []Drip{{Id: 5, Text: ""}, {Id: 7, Text: ""}, {Id: 9, Text: ""}}, expected: 10},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("drips=%v", tc.drips)
		t.Run(name, func(t *testing.T) {
			result := GetNextId(tc.drips)
			msg := fmt.Sprintf("%s,result=%d,expected=%d", name, result, tc.expected)
			assert.Assert(result == tc.expected, msg)
		})
	}
}

func TestMaxDripIdWidth(t *testing.T) {
	type testMaxIdWidth struct {
		drips    []Drip
		expected int
	}

	testCases := []testMaxIdWidth{
		{drips: []Drip{}, expected: 0},
		{drips: []Drip{{Id: 0, Text: ""}}, expected: 1},
		{drips: []Drip{{Id: 5, Text: ""}}, expected: 1},
		{drips: []Drip{{Id: 10, Text: ""}}, expected: 2},
		{drips: []Drip{{Id: 123, Text: ""}}, expected: 3},
		{drips: []Drip{{Id: 9999, Text: ""}}, expected: 4},
		{drips: []Drip{{Id: -1, Text: ""}}, expected: 2},
		{drips: []Drip{{Id: -100, Text: ""}}, expected: 4},
		{drips: []Drip{{Id: 9, Text: ""}, {Id: 5, Text: ""}}, expected: 1},
		{drips: []Drip{{Id: 100, Text: ""}, {Id: 5, Text: ""}}, expected: 3},
		{drips: []Drip{{Id: 7, Text: ""}, {Id: 7, Text: ""}}, expected: 1},
		{drips: []Drip{{Id: -5, Text: ""}, {Id: -3, Text: ""}, {Id: -10, Text: ""}}, expected: 3},
		{drips: []Drip{{Id: -5, Text: ""}, {Id: 0, Text: ""}, {Id: 3, Text: ""}}, expected: 2},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("drips=%v", tc.drips)
		t.Run(name, func(t *testing.T) {
			result := MaxDripIdWidth(tc.drips)
			msg := fmt.Sprintf("%s,result=%d,expected=%d", name, result, tc.expected)
			assert.Assert(result == tc.expected, msg)
		})
	}
}

func TestFindIndexFromId(t *testing.T) {
	type testFindIndexFromId struct {
		name     string
		drips    []Drip
		id       int
		expected int
	}

	testCases := []testFindIndexFromId{
		{name: "empty", drips: []Drip{}, id: 1, expected: 0},
		{name: "single_found", drips: []Drip{{Id: 1, Text: ""}}, id: 1, expected: 0},
		{name: "single_not_found_less", drips: []Drip{{Id: 5, Text: ""}}, id: 3, expected: 0},
		{name: "single_not_found_greater", drips: []Drip{{Id: 5, Text: ""}}, id: 10, expected: 1},
		{name: "multiple_found_first", drips: []Drip{{Id: 1, Text: ""}, {Id: 11, Text: ""}}, id: 1, expected: 0},
		{name: "multiple_found_last", drips: []Drip{{Id: 1, Text: ""}, {Id: 11, Text: ""}}, id: 11, expected: 1},
		{name: "multiple_found_middle", drips: []Drip{{Id: 1, Text: ""}, {Id: 5, Text: ""}, {Id: 11, Text: ""}}, id: 5, expected: 1},
		{name: "not_found_between", drips: []Drip{{Id: 1, Text: ""}, {Id: 5, Text: ""}, {Id: 11, Text: ""}}, id: 7, expected: 2},
		{name: "not_found_before_all", drips: []Drip{{Id: 1, Text: ""}, {Id: 5, Text: ""}, {Id: 11, Text: ""}}, id: -1, expected: 0},
		{name: "not_found_after_all", drips: []Drip{{Id: 1, Text: ""}, {Id: 5, Text: ""}, {Id: 11, Text: ""}}, id: 100, expected: 3},
		{name: "large_slice", drips: []Drip{{Id: 1, Text: ""}, {Id: 11, Text: ""}, {Id: 111, Text: ""}, {Id: 1111, Text: ""}}, id: 1111, expected: 3},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%s", tc.name)
		t.Run(name, func(t *testing.T) {
			result := FindIndexFromId(tc.drips, tc.id)
			msg := fmt.Sprintf("%s,result=%d,expected=%d", name, result, tc.expected)
			assert.Assert(result == tc.expected, msg)
		})
	}
}

func TestSortDripsById(t *testing.T) {
	type testSortDripsById struct {
		name     string
		input    []Drip
		expected []int // expected order of Ids after sort
	}

	testCases := []testSortDripsById{
		{
			name:     "empty",
			input:    []Drip{},
			expected: []int{},
		},
		{
			name:     "single",
			input:    []Drip{{Id: 5, Text: "Drip 5"}},
			expected: []int{5},
		},
		{
			name:     "already sorted",
			input:    []Drip{{Id: 1, Text: "A"}, {Id: 2, Text: "B"}, {Id: 3, Text: "C"}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "reverse sorted",
			input:    []Drip{{Id: 3, Text: "C"}, {Id: 2, Text: "B"}, {Id: 1, Text: "A"}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "duplicate ids",
			input:    []Drip{{Id: 2, Text: "B1"}, {Id: 1, Text: "A"}, {Id: 2, Text: "B2"}},
			expected: []int{1, 2, 2},
		},
		{
			name:     "negative ids",
			input:    []Drip{{Id: -5, Text: "Neg5"}, {Id: -1, Text: "Neg1"}, {Id: -10, Text: "Neg10"}},
			expected: []int{-10, -5, -1},
		},
		{
			name:     "mixed positive_negative",
			input:    []Drip{{Id: 5, Text: "Pos5"}, {Id: -3, Text: "Neg3"}, {Id: 0, Text: "Zero"}},
			expected: []int{-3, 0, 5},
		},
		{
			name:     "duplicate ids reverse_order",
			input:    []Drip{{Id: 3, Text: "C"}, {Id: 1, Text: "A"}, {Id: 3, Text: "C2"}, {Id: 2, Text: "B"}},
			expected: []int{1, 2, 3, 3},
		},
		{
			name:     "with_duplicate_text",
			input:    []Drip{{Id: 10, Text: "Same text"}, {Id: 5, Text: "Same text"}, {Id: 15, Text: "Different"}},
			expected: []int{5, 10, 15},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// make a copy to avoid mutating test case
			drips := make([]Drip, len(tc.input))
			copy(drips, tc.input)

			SortDripsById(drips)

			if len(drips) != len(tc.expected) {
				assert.Assert(false, fmt.Sprintf("%s: length mismatch got %d want %d", tc.name, len(drips), len(tc.expected)))
				return
			}
			for i, id := range tc.expected {
				if drips[i].Id != id {
					assert.Assert(false, fmt.Sprintf("%s: at index %d got Id %d want %d", tc.name, i, drips[i].Id, id))
					return
				}
			}
		})
	}
}

func TestUpdateDripsIds(t *testing.T) {
	type testUpdateDripsIds struct {
		name     string
		drips    []Drip
		startIdx int
		idToAdd  int
		dripsLen int // expected length after update
		// Check which specific drip IDs changed
		elementIds []int // expected IDs at specific indices after update
	}

	testCases := []testUpdateDripsIds{
		{
			name:       "basic_update_start_at_beginning",
			drips:      []Drip{{Id: 1, Text: "A"}, {Id: 2, Text: "B"}, {Id: 3, Text: "C"}, {Id: 4, Text: "D"}},
			startIdx:   0,
			idToAdd:    2,
			dripsLen:   4,
			elementIds: []int{1, 3, 4, 5}, // [1, 2->3, 3->4, 4->5]
		},
		{
			name:       "basic_update_middle",
			drips:      []Drip{{Id: 1, Text: "A"}, {Id: 2, Text: "B"}, {Id: 3, Text: "C"}, {Id: 4, Text: "D"}},
			startIdx:   1,
			idToAdd:    3,
			dripsLen:   4,
			elementIds: []int{1, 2, 4, 5}, // [1, 2, 3->4, 4->5]
		},
		{
			name:       "update_last_element",
			drips:      []Drip{{Id: 1, Text: "A"}, {Id: 2, Text: "B"}, {Id: 3, Text: "C"}, {Id: 4, Text: "D"}},
			startIdx:   3,
			idToAdd:    4,
			dripsLen:   4,
			elementIds: []int{1, 2, 3, 5}, // [1, 2, 3, 4->5]
		},
		{
			name:       "no_match_found",
			drips:      []Drip{{Id: 1, Text: "A"}, {Id: 2, Text: "B"}, {Id: 3, Text: "C"}},
			startIdx:   0,
			idToAdd:    5, // ID that doesn't exist
			dripsLen:   3,
			elementIds: []int{1, 2, 3}, // unchanged
		},
		{
			name:       "empty_slice",
			drips:      []Drip{},
			startIdx:   0,
			idToAdd:    1,
			dripsLen:   0,
			elementIds: []int{},
		},
		{
			name:       "update_duplicate_ids",
			drips:      []Drip{{Id: 1, Text: "A"}, {Id: 2, Text: "B"}, {Id: 2, Text: "B2"}, {Id: 3, Text: "C"}},
			startIdx:   0,
			idToAdd:    2,
			dripsLen:   4,
			elementIds: []int{1, 3, 2, 4}, // [1, 2->3, 2 unchanged, 3->4]
		},
		{
			name:       "id_less_than_start",
			drips:      []Drip{{Id: 1, Text: "A"}, {Id: 2, Text: "B"}, {Id: 3, Text: "C"}},
			startIdx:   2,
			idToAdd:    1, // ID before startIdx, should not be updated
			dripsLen:   3,
			elementIds: []int{1, 2, 3}, // unchanged because search starts at index 2
		},
		{
			name:       "update_multiple_consecutive_ids",
			drips:      []Drip{{Id: 1, Text: "A"}, {Id: 2, Text: "B"}, {Id: 3, Text: "C"}, {Id: 4, Text: "D"}, {Id: 5, Text: "E"}},
			startIdx:   1,
			idToAdd:    3,
			dripsLen:   5,
			elementIds: []int{1, 2, 4, 5, 6}, // [1, 2, 3->4, 4->5, 5->6]
		},
		{
			name:       "update_first_element",
			drips:      []Drip{{Id: 1, Text: "A"}, {Id: 2, Text: "B"}, {Id: 3, Text: "C"}},
			startIdx:   0,
			idToAdd:    1,
			dripsLen:   3,
			elementIds: []int{2, 3, 4}, // [1->2, 2->3, 3->4]
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// make a copy to avoid mutating test case
			drips := make([]Drip, len(tc.drips))
			copy(drips, tc.drips)

			UpdateDripsIds(drips, tc.startIdx, tc.idToAdd)

			// Check length
			if len(drips) != tc.dripsLen {
				assert.Assert(false, fmt.Sprintf("%s: expected %d drips, got %d", tc.name, tc.dripsLen, len(drips)))
			}

			// Check specific element IDs if provided
			if tc.elementIds != nil {
				if len(drips) != len(tc.elementIds) {
					assert.Assert(false, fmt.Sprintf("%s: expected %d elements, got %d", tc.name, len(tc.elementIds), len(drips)))
					return
				}

				for i, expectedId := range tc.elementIds {
					if drips[i].Id != expectedId {
						assert.Assert(false, fmt.Sprintf("%s: at index %d expected Id %d, got %d", tc.name, i, expectedId, drips[i].Id))
						return
					}
				}
			}

			// Additional check: verify original Ids are preserved where they shouldn't change
			for i, originalDrip := range tc.drips {
				if i < tc.startIdx {
					// Elements before startIdx should never change
					if drips[i].Id != originalDrip.Id {
						assert.Assert(false, fmt.Sprintf("%s: element at index %d before startIdx should not change: expected %d, got %d", tc.name, i, originalDrip.Id, drips[i].Id))
					}
				}
			}
		})
	}
}
