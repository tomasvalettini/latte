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

//func TestBlendIdentifierIsInvalid(t *testing.T) {
//	listOfBlendIdentifier := getBlendIdentifierList(-1, false)
//
//	for _, bi := range listOfBlendIdentifier {
//		assert.Assert(bi.IsValid() == false, "BlendIdentifier should not be valid")
//	}
//}
//
//
//func TestBlendIdentifierIsValid(t *testing.T) {
//	listOfBlendIdentifier := getBlendIdentifierList(1, true)
//
//	for _, bi := range listOfBlendIdentifier {
//		assert.Assert(bi.IsValid() == true, "BlendIdentifier should be valid")
//	}
//}
//
//func TestBlendIdentifierIsIdInvalid(t *testing.T) {
//	listOfBlendIdentifier := getBlendIdentifierList(-1, true)
//
//	for _, bi := range listOfBlendIdentifier {
//		assert.Assert(bi.IsIdValid() == false, "BlendIdentifier id should not be valid")
//	}
//}
//
//func TestBlendIdentifierIsIdValid(t *testing.T) {
//	listOfBlendIdentifier := getBlendIdentifierList(1, true)
//
//	for _, bi := range listOfBlendIdentifier {
//		assert.Assert(bi.IsValid() == true, "BlendIdentifier id should be valid")
//	}
//}
//
//func TestBlendIdentifierIsTitleInvalid(t *testing.T) {
//	listOfBlendIdentifier := getBlendIdentifierList(-1, false)
//
//	for _, bi := range listOfBlendIdentifier {
//		assert.Assert(bi.IsTitleValid() == false, "BlendIdentifier title should not be valid")
//	}
//}
//
//func TestBlendIdentifierIsTitleValid(t *testing.T) {
//	listOfBlendIdentifier := getBlendIdentifierList(-1, true)
//
//	for _, bi := range listOfBlendIdentifier {
//		assert.Assert(bi.IsTitleValid() == true, "BlendIdentifier title should be valid")
//	}
//}
//
//func getBlendIdentifierList(mul int, incTitle bool) []BlendIdentifier {
//	listOfBlendIdentifier := []BlendIdentifier{}
//
//	for i:= 1; i < 10; i++ {
//		title := ""
//
//		if incTitle {
//			title = "test title " + strconv.Itoa(i)
//		}
//
//		blendIdentifier := BlendIdentifier{
//			Id: i * mul,
//			Title: title,
//		}
//
//		listOfBlendIdentifier = append(listOfBlendIdentifier, blendIdentifier)
//	}
//
//	return listOfBlendIdentifier
//}
