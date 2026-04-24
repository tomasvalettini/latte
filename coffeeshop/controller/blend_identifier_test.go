package controller

import (
	"fmt"
	"testing"

	"github.com/tomasvalettini/latte/assert"
)

type TestBlendIdentifier struct {
	bi       BlendIdentifier
	expected bool
}

// TestIsValid checks if IsValid returns true when both Id and Title are valid.
func TestIsValid(t *testing.T) {
	testCases := []TestBlendIdentifier{
		{BlendIdentifier{Id: 0, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 1, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 2, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 3, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 4, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 1000, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 0, Title: ""}, true},
		{BlendIdentifier{Id: 1, Title: ""}, true},
		{BlendIdentifier{Id: 2, Title: ""}, true},
		{BlendIdentifier{Id: 3, Title: ""}, true},
		{BlendIdentifier{Id: 4, Title: ""}, true},
		{BlendIdentifier{Id: 1000, Title: ""}, true},
		{BlendIdentifier{Id: -1, Title: "Coffee"}, true},
		{BlendIdentifier{Id: -2, Title: "Coffee"}, true},
		{BlendIdentifier{Id: -3, Title: "Coffee"}, true},
		{BlendIdentifier{Id: -4, Title: "Coffee"}, true},
		{BlendIdentifier{Id: -5, Title: "Coffee"}, true},
		{BlendIdentifier{Id: -1000, Title: "Coffee"}, true},
		{BlendIdentifier{Id: -1, Title: ""}, false},
		{BlendIdentifier{Id: -2, Title: ""}, false},
		{BlendIdentifier{Id: -3, Title: ""}, false},
		{BlendIdentifier{Id: -4, Title: ""}, false},
		{BlendIdentifier{Id: -5, Title: ""}, false},
		{BlendIdentifier{Id: -1000, Title: ""}, false},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("Id=%d,Title=%s", tc.bi.Id, tc.bi.Title)
		t.Run(name, func(t *testing.T) {
			result := tc.bi.IsValid()
			msg := fmt.Sprintf("%s,result=%t,expected=%t", name, result, tc.expected)
			assert.Assert(result == tc.expected, msg)
		})
	}
}

// TestIsIdValid checks if IsIdValid returns true when Id is valid.
func TestIsIdValid(t *testing.T) {
	testCases := []TestBlendIdentifier{
		{BlendIdentifier{Id: 1, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 0, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 1000, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 1, Title: ""}, true},
		{BlendIdentifier{Id: 0, Title: ""}, true},
		{BlendIdentifier{Id: 1000, Title: ""}, true},
		{BlendIdentifier{Id: -1, Title: "Coffee"}, false},
		{BlendIdentifier{Id: -1000, Title: "Coffee"}, false},
		{BlendIdentifier{Id: -1, Title: ""}, false},
		{BlendIdentifier{Id: -1000, Title: ""}, false},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("Id=%d", tc.bi.Id)
		t.Run(name, func(t *testing.T) {
			result := tc.bi.IsIdValid()
			msg := fmt.Sprintf("%s,result=%t,expected=%t", name, result, tc.expected)
			assert.Assert(result == tc.expected, msg)
		})
	}
}

// TestIsTitleValid checks if the IsTitleValid method returns true when Title is valid.
func TestIsTitleValid(t *testing.T) {
	testCases := []TestBlendIdentifier{
		{BlendIdentifier{Id: 1, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 0, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 1000, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 1, Title: ""}, false},
		{BlendIdentifier{Id: 0, Title: ""}, false},
		{BlendIdentifier{Id: 1000, Title: ""}, false},
		{BlendIdentifier{Id: -1, Title: "Coffee"}, true},
		{BlendIdentifier{Id: -1000, Title: "Coffee"}, true},
		{BlendIdentifier{Id: -1, Title: ""}, false},
		{BlendIdentifier{Id: -1000, Title: ""}, false},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("Id=%d", tc.bi.Id)
		t.Run(name, func(t *testing.T) {
			result := tc.bi.IsTitleValid()
			msg := fmt.Sprintf("%s,result=%t,expected=%t", name, result, tc.expected)
			assert.Assert(result == tc.expected, msg)
		})
	}
}

// TestValidate checks if Validate returns a non-nil pointer when valid and nil when invalid.
func TestValidate(t *testing.T) {
	testCases := []struct {
		bi       BlendIdentifier
		expected bool // true if should return non-nil, false if should return nil
	}{
		// Valid cases: Id >= 0 or Title is not empty
		{BlendIdentifier{Id: 0, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 1, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 1000, Title: "Coffee"}, true},
		{BlendIdentifier{Id: 0, Title: ""}, true},
		{BlendIdentifier{Id: 1, Title: ""}, true},
		{BlendIdentifier{Id: 1000, Title: ""}, true},
		{BlendIdentifier{Id: -1, Title: "Coffee"}, true},
		{BlendIdentifier{Id: -2, Title: "Espresso"}, true},
		{BlendIdentifier{Id: -1000, Title: "Latte"}, true},
		// Invalid cases: Id < 0 AND Title is empty
		{BlendIdentifier{Id: -1, Title: ""}, false},
		{BlendIdentifier{Id: -2, Title: ""}, false},
		{BlendIdentifier{Id: -1000, Title: ""}, false},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("Id=%d,Title=%s", tc.bi.Id, tc.bi.Title)
		t.Run(name, func(t *testing.T) {
			result := tc.bi.Validate()
			isNonNil := result != nil
			msg := fmt.Sprintf("%s,result=%v,expected=%t", name, result, tc.expected)
			assert.Assert(isNonNil == tc.expected, msg)

			// Verify that the returned pointer points to the correct BlendIdentifier
			if result != nil {
				assert.Assert(result.Id == tc.bi.Id && result.Title == tc.bi.Title,
					fmt.Sprintf("%s,returned pointer has incorrect values", name))
			}
		})
	}
}
