package controller

import (
	"fmt"
	"os"
	"testing"

	"github.com/tomasvalettini/latte/assert"
	carafepath "github.com/tomasvalettini/latte/coffeeshop/data/data-source/path"
	datamodel "github.com/tomasvalettini/latte/coffeeshop/data/model"
)

// Test adding to nil identifier defaults to House Blend
func TestAddToBlends_WithNilIdentifier(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Add drips to nil (should default to house blend)
	tc.AddToBlends(nil, "test drip 1")
	tc.AddToBlends(nil, "test drip 2")
	tc.AddToBlends(nil, "test drip 3")

	blends := tc.dataSource.Load()
	houseBlend := findBlendByTitle(blends, HOUSE_BLEND_TITLE)

	dripsLen := len(houseBlend.Drips)
	dripText := houseBlend.Drips[0].Text

	performTestChecks(
		map[string]bool{
			"House blend not found": houseBlend == nil,
			fmt.Sprintf("Expected 3 drips in house blend, got %d", dripsLen):           dripsLen != 3,
			fmt.Sprintf("Expected first drip to be 'test drip 1', got '%s'", dripText): dripText != "test drip 1",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test rejecting empty drip text
func TestAddToBlends_WithEmptyDripText(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Try to add empty drip
	tc.AddToBlends(nil, "")

	blends := tc.dataSource.Load()

	performTestChecks(
		map[string]bool{
			"Empty drip should not be added": len(blends) > 0 && len(blends[0].Drips) > 0,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test creating a new blend
func TestAddToBlends_CreatesNewBlend(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Add drips to a new blend
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "shot 1")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "shot 2")

	blends := tc.dataSource.Load()
	espressoBlend := findBlendByTitle(blends, "Espresso")

	dripLen := len(espressoBlend.Drips)

	performTestChecks(
		map[string]bool{
			"Espresso blend not found": espressoBlend == nil,
			fmt.Sprintf("Expected 2 drips in Espresso blend, got %d", dripLen): dripLen != 2,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test adding drips to existing blend by title
func TestAddToBlends_AddsToExistingBlendByTitle(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Create a blend
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Cappuccino"}, "drip 1")

	// Add more drips using the same blend identifier
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Cappuccino"}, "drip 2")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Cappuccino"}, "drip 3")

	blends := tc.dataSource.Load()
	cappuccinoBlend := findBlendByTitle(blends, "Cappuccino")

	// Should only have one blend with this title
	blendCount := 0
	for _, blend := range blends {
		if blend.Title == "Cappuccino" {
			blendCount++
		}
	}

	dripLen := len(cappuccinoBlend.Drips)

	performTestChecks(
		map[string]bool{
			"Cappuccino blend not found":                                         cappuccinoBlend == nil,
			fmt.Sprintf("Expected 1 Cappuccino blend, found %d", blendCount):     blendCount != 1,
			fmt.Sprintf("Expected 3 drips in Cappuccino blend, got %d", dripLen): dripLen != 3,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test auto-generating drip IDs
func TestAddToBlends_AutoGeneratesDripIds(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Add drips to a blend
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Test Blend"}, "drip 1")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Test Blend"}, "drip 2")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Test Blend"}, "drip 3")

	blends := tc.dataSource.Load()
	testBlend := findBlendByTitle(blends, "Test Blend")

	performTestChecks(
		map[string]bool{
			"Test Blend not found": testBlend == nil,
		},
	)

	// Verify that drip IDs are auto-incremented and unique
	expectedIds := []int{0, 1, 2}
	for i, expectedId := range expectedIds {
		dripId := testBlend.Drips[i].Id
		performTestChecks(
			map[string]bool{
				fmt.Sprintf("Expected drip %d to have id %d, got %d", i, expectedId, dripId): dripId != expectedId,
			},
		)
	}

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test listing blends with empty data
func TestListBlends_WithNoBlends(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Don't add any blends, just call ListBlends
	// This should print "Nothing to show yet!" but we verify by checking data source
	blends := tc.dataSource.Load()

	performTestChecks(
		map[string]bool{
			"Expected no blends, but some were found": len(blends) != 0,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test listing all blends with nil identifier
func TestListBlends_WithNilIdentifier(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add multiple blends
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "shot 1")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Cappuccino"}, "milk 1")
	tc.AddToBlends(nil, "house drip")

	// Call ListBlends with nil to list all
	tc.ListBlends(nil)

	blends := tc.dataSource.Load()

	blendsLen := len(blends)

	performTestChecks(
		map[string]bool{
			fmt.Sprintf("Expected 2 blends, got %d", blendsLen): blendsLen < 2,
		},
	)

	// Verify Espresso exists
	hasEspresso := findBlendByTitle(blends, "Espresso") != nil
	hasCapuccino := findBlendByTitle(blends, "Cappuccino") != nil

	performTestChecks(
		map[string]bool{
			"Espresso blend not found":   !hasEspresso,
			"Cappuccino blend not found": !hasCapuccino,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test listing specific blend by title
func TestListBlends_ByTitle(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add multiple blends
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "shot 1")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "shot 2")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Cappuccino"}, "milk 1")

	// Call ListBlends for specific blend
	tc.ListBlends(&BlendIdentifier{Id: -1, Title: "Espresso"})

	// Verify the blend exists and has the right drips
	blends := tc.dataSource.Load()
	espressoBlend := findBlendByTitle(blends, "Espresso")

	dripCount := len(espressoBlend.Drips)

	performTestChecks(
		map[string]bool{
			"Espresso blend not found":                                     espressoBlend == nil,
			fmt.Sprintf("Expected 2 drips in Espresso, got %d", dripCount): dripCount != 2,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test listing specific blend by ID
func TestListBlends_ById(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add blends
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "shot 1")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Cappuccino"}, "milk 1")

	blends := tc.dataSource.Load()
	espressoBlend := findBlendByTitle(blends, "Espresso")
	var espressoId int
	if espressoBlend != nil {
		espressoId = espressoBlend.Id
	}

	performTestChecks(
		map[string]bool{
			"Espresso blend not found": espressoBlend == nil,
		},
	)

	// Call ListBlends with ID
	tc.ListBlends(&BlendIdentifier{Id: espressoId, Title: ""})

	// Verify the blend still exists
	blends = tc.dataSource.Load()
	var foundBlend *datamodel.Blend
	for i := range blends {
		if blends[i].Id == espressoId {
			foundBlend = &blends[i]
			break
		}
	}
	// Note: searching by ID, keeping the loop as there's no title for ID-based lookup

	performTestChecks(
		map[string]bool{
			"Blend with given ID not found": foundBlend == nil,
			fmt.Sprintf("Expected blend title to be 'Espresso', got '%s'", foundBlend.Title): foundBlend.Title != "Espresso",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test listing with invalid blend identifier
func TestListBlends_WithInvalidIdentifier(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "shot 1")

	// Call ListBlends with invalid identifier (both id and title empty/invalid)
	tc.ListBlends(&BlendIdentifier{Id: -999, Title: ""})

	// Just verify that blends still exist (no side effects)
	blends := tc.dataSource.Load()
	blendsLen := len(blends)

	performTestChecks(
		map[string]bool{
			fmt.Sprintf("Expected 1 blend to remain, got %d", blendsLen): blendsLen != 1,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test deleting a drip from a blend by dripId
func TestDeleteFromBlends_DeleteWholeBlend(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend with multiple drips
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Test Blend"}, "drip 1")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Test Blend"}, "drip 2")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Test Blend"}, "drip 3")

	blends := tc.dataSource.Load()

	performTestChecks(
		map[string]bool{
			"Expected at least one blend after setup": len(blends) <= 0,
		},
	)

	// Mock stdin for the confirmation prompt
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("yes\n")
	w.Close()

	// Delete drip with id 0 (negative dripId means delete specific drip)
	tc.DeleteFromBlends(&BlendIdentifier{Id: -1, Title: "Test Blend"}, -1)

	// Restore stdin
	os.Stdin = oldStdin

	// Verify the drip was deleted by loading the blends
	blends = tc.dataSource.Load()
	testBlend := findBlendByTitle(blends, "Test Blend")

	performTestChecks(
		map[string]bool{
			"Test blend should not found after deletion": testBlend != nil,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test deleting a drip with negative id from a blend
func TestDeleteFromBlends_DeleteDripWithNegativeId(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add blends and drips
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "shot 1")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "shot 2")

	// Delete the first drip (id 0)
	tc.DeleteFromBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, 0)

	blends := tc.dataSource.Load()
	espressoBlend := findBlendByTitle(blends, "Espresso")

	dripCount := len(espressoBlend.Drips)
	remainingDripText := ""
	if dripCount > 0 {
		remainingDripText = espressoBlend.Drips[0].Text
	}

	performTestChecks(
		map[string]bool{
			"Espresso blend not found":                                                         espressoBlend == nil,
			fmt.Sprintf("Expected 1 drip after deletion, got %d", dripCount):                   dripCount != 1,
			fmt.Sprintf("Expected remaining drip to be 'shot 2', got '%s'", remainingDripText): remainingDripText != "shot 2",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test deleting a drip that doesn't exist
func TestDeleteFromBlends_DripNotFound(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend with a drip
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Cappuccino"}, "drip 1")

	// Try to delete a drip that doesn't exist (id 999)
	tc.DeleteFromBlends(&BlendIdentifier{Id: -1, Title: "Cappuccino"}, -999)

	// Verify the blend and drip still exist
	blends := tc.dataSource.Load()
	cappuccinoBlend := findBlendByTitle(blends, "Cappuccino")

	dripCount := len(cappuccinoBlend.Drips)

	performTestChecks(
		map[string]bool{
			"Cappuccino blend not found":                                cappuccinoBlend == nil,
			fmt.Sprintf("Expected 1 drip to remain, got %d", dripCount): dripCount != 1,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test deleting from a blend that doesn't exist
func TestDeleteFromBlends_BlendNotFound(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Real Blend"}, "drip 1")

	// Try to delete from a non-existent blend (use 0 as dripId to delete specific drip)
	tc.DeleteFromBlends(&BlendIdentifier{Id: -1, Title: "Non-existent Blend"}, 0)

	// Verify the original blend is unchanged
	blends := tc.dataSource.Load()

	performTestChecks(
		map[string]bool{
			fmt.Sprintf("Expected 1 blend, got %d", len(blends)): len(blends) != 1,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test deleting with nil BlendIdentifier defaults to House Blend
func TestDeleteFromBlends_DefaultsToHouseBlend(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add drips to the default house blend
	tc.AddToBlends(nil, "coffee 1")
	tc.AddToBlends(nil, "coffee 2")
	tc.AddToBlends(nil, "coffee 3")

	// Mock stdin for the confirmation prompt
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("yes\n")
	w.Close()

	// Delete from nil identifier (should default to house blend) - use -1 to delete whole blend
	tc.DeleteFromBlends(nil, -1)

	// Restore stdin
	os.Stdin = oldStdin

	blends := tc.dataSource.Load()
	houseBlend := findBlendByTitle(blends, HOUSE_BLEND_TITLE)

	performTestChecks(
		map[string]bool{
			"House blend should be deleted": houseBlend != nil,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test deleting multiple drips from the same blend
func TestDeleteFromBlends_DeleteMultipleDrips(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add multiple drips
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Multi Blend"}, "drip 1")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Multi Blend"}, "drip 2")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Multi Blend"}, "drip 3")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Multi Blend"}, "drip 4")

	// Delete first drip (id 0)
	tc.DeleteFromBlends(&BlendIdentifier{Id: -1, Title: "Multi Blend"}, 0)

	blends := tc.dataSource.Load()
	multiBlend := findBlendByTitle(blends, "Multi Blend")

	dripCountAfterFirst := len(multiBlend.Drips)

	performTestChecks(
		map[string]bool{
			"Multi blend not found": multiBlend == nil,
			fmt.Sprintf("Expected 3 drips after first deletion, got %d", dripCountAfterFirst): dripCountAfterFirst != 3,
		},
	)

	// Delete another drip (id 1 from the remaining)
	tc.DeleteFromBlends(&BlendIdentifier{Id: -1, Title: "Multi Blend"}, 1)

	blends = tc.dataSource.Load()
	multiBlend = findBlendByTitle(blends, "Multi Blend")

	dripCountAfterSecond := len(multiBlend.Drips)

	performTestChecks(
		map[string]bool{
			fmt.Sprintf("Expected 2 drips after second deletion, got %d", dripCountAfterSecond): dripCountAfterSecond != 2,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test deleting from empty blend
func TestDeleteFromBlends_EmptyBlend(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Create a blend with no drips (indirectly via data source manipulation)
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Empty Blend"}, "temp drip")

	// Mock stdin for the first confirmation prompt to delete the blend
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("yes\n")
	w.Close()

	tc.DeleteFromBlends(&BlendIdentifier{Id: -1, Title: "Empty Blend"}, -1)

	// Restore stdin
	os.Stdin = oldStdin

	blends := tc.dataSource.Load()
	emptyBlend := findBlendByTitle(blends, "Empty Blend")

	performTestChecks(
		map[string]bool{
			"Empty blend should be deleted": emptyBlend != nil,
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test successfully updating a drip text in a blend
func TestUpdateDripInBlend_SuccessfulUpdate(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend with multiple drips
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Update Test Blend"}, "original drip 1")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Update Test Blend"}, "original drip 2")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Update Test Blend"}, "original drip 3")

	// Update the second drip (id 1)
	tc.UpdateDripInBlend(&BlendIdentifier{Id: -1, Title: "Update Test Blend"}, 1, "updated drip 2")

	blends := tc.dataSource.Load()
	testBlend := findBlendByTitle(blends, "Update Test Blend")

	updatedDripText := testBlend.Drips[1].Text
	dripsLen := len(testBlend.Drips)

	performTestChecks(
		map[string]bool{
			"Update Test Blend not found":                                                          testBlend == nil,
			fmt.Sprintf("Expected 3 drips to remain, got %d", dripsLen):                            dripsLen != 3,
			fmt.Sprintf("Expected updated drip to be 'updated drip 2', got '%s'", updatedDripText): updatedDripText != "updated drip 2",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test updating drip with invalid (negative) drip ID
func TestUpdateDripInBlend_WithInvalidDripId(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend with a drip
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Negative ID Blend"}, "test drip")

	// Try to update with negative drip ID and non-empty text (should fail)
	tc.UpdateDripInBlend(&BlendIdentifier{Id: -1, Title: "Negative ID Blend"}, -1, "new text")

	// Verify the drip was not changed
	blends := tc.dataSource.Load()
	negativeIdBlend := findBlendByTitle(blends, "Negative ID Blend")

	dripText := negativeIdBlend.Drips[0].Text

	performTestChecks(
		map[string]bool{
			"Negative ID Blend not found": negativeIdBlend == nil,
			fmt.Sprintf("Expected drip text to remain 'test drip', got '%s'", dripText): dripText != "test drip",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test updating blend title when dripId < 0 and dripText is empty
func TestUpdateDripInBlend_UpdatesBlendTitle(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Old Title"}, "drip 1")

	// Get the blend ID to verify title update
	blends := tc.dataSource.Load()
	oldBlend := findBlendByTitle(blends, "Old Title")
	blendId := oldBlend.Id

	fmt.Println(blendId)

	// Update blend title using negative dripId and empty dripText
	tc.UpdateDripInBlend(&BlendIdentifier{Id: blendId, Title: "New Title"}, -1, "")

	// Verify the blend title was updated
	blends = tc.dataSource.Load()
	var updatedBlend *datamodel.Blend = nil
	for i := range blends {
		fmt.Println(blends[i])
		if blends[i].Id == blendId {
			updatedBlend = &blends[i]
			break
		}
	}

	performTestChecks(
		map[string]bool{
			"Blend not found": updatedBlend == nil,
			fmt.Sprintf("Expected blend title to be 'New Title', got '%s'", updatedBlend.Title): updatedBlend.Title != "New Title",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test updating blend title with non-existent blend
func TestUpdateDripInBlend_UpdateTitleBlendNotFound(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a real blend
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Real Blend"}, "drip 1")

	// Try to update title of non-existent blend
	tc.UpdateDripInBlend(&BlendIdentifier{Id: -1, Title: "Non-existent Blend"}, -1, "")

	// Verify real blend is unchanged
	blends := tc.dataSource.Load()
	realBlend := findBlendByTitle(blends, "Real Blend")

	performTestChecks(
		map[string]bool{
			"Real Blend not found": realBlend == nil,
			fmt.Sprintf("Expected blend title to remain 'Real Blend', got '%s'", realBlend.Title): realBlend.Title != "Real Blend",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test updating drip with empty text
func TestUpdateDripInBlend_WithEmptyText(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend with a drip
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Empty Text Blend"}, "original drip")

	// Try to update with empty text
	tc.UpdateDripInBlend(&BlendIdentifier{Id: -1, Title: "Empty Text Blend"}, 0, "")

	// Verify the drip was not changed
	blends := tc.dataSource.Load()
	emptyTextBlend := findBlendByTitle(blends, "Empty Text Blend")

	dripText := emptyTextBlend.Drips[0].Text

	performTestChecks(
		map[string]bool{
			"Empty Text Blend not found": emptyTextBlend == nil,
			fmt.Sprintf("Expected drip text to remain 'original drip', got '%s'", dripText): dripText != "original drip",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test updating drip with nil blend identifier
func TestUpdateDripInBlend_WithNilIdentifier(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a drip to house blend
	tc.AddToBlends(nil, "house drip")

	// Try to update with nil identifier (should reject)
	tc.UpdateDripInBlend(nil, 0, "new text")

	// Verify the drip was not changed
	blends := tc.dataSource.Load()
	houseBlend := findBlendByTitle(blends, HOUSE_BLEND_TITLE)

	dripText := houseBlend.Drips[0].Text

	performTestChecks(
		map[string]bool{
			"House Blend not found": houseBlend == nil,
			fmt.Sprintf("Expected drip text to remain 'house drip', got '%s'", dripText): dripText != "house drip",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test updating drip in non-existent blend
func TestUpdateDripInBlend_BlendNotFound(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a real blend
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Real Blend"}, "drip 1")

	// Try to update drip in non-existent blend
	tc.UpdateDripInBlend(&BlendIdentifier{Id: -1, Title: "Non-existent Blend"}, 0, "new text")

	// Verify real blend is unchanged
	blends := tc.dataSource.Load()
	realBlend := findBlendByTitle(blends, "Real Blend")

	dripText := realBlend.Drips[0].Text

	performTestChecks(
		map[string]bool{
			"Real Blend not found": realBlend == nil,
			fmt.Sprintf("Expected drip text to remain 'drip 1', got '%s'", dripText): dripText != "drip 1",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test updating drip with non-existent drip ID
func TestUpdateDripInBlend_DripNotFound(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend with a drip (id 0)
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Not Found Blend"}, "drip 1")

	// Try to update drip with non-existent ID
	tc.UpdateDripInBlend(&BlendIdentifier{Id: -1, Title: "Not Found Blend"}, 999, "new text")

	// Verify the drip was not changed
	blends := tc.dataSource.Load()
	notFoundBlend := findBlendByTitle(blends, "Not Found Blend")

	dripText := notFoundBlend.Drips[0].Text

	performTestChecks(
		map[string]bool{
			"Not Found Blend not found": notFoundBlend == nil,
			fmt.Sprintf("Expected drip text to remain 'drip 1', got '%s'", dripText): dripText != "drip 1",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test updating drip by blend ID (not title)
func TestUpdateDripInBlend_ByBlendId(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend with a drip
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "ID Based Blend"}, "original")

	// Get the blend ID
	blends := tc.dataSource.Load()
	idBasedBlend := findBlendByTitle(blends, "ID Based Blend")
	blendId := idBasedBlend.Id

	// Update using blend ID instead of title
	tc.UpdateDripInBlend(&BlendIdentifier{Id: blendId, Title: ""}, 0, "updated by id")

	// Verify the drip was updated
	blends = tc.dataSource.Load()
	var updatedBlend *datamodel.Blend
	for i := range blends {
		if blends[i].Id == blendId {
			updatedBlend = &blends[i]
			break
		}
	}

	dripText := updatedBlend.Drips[0].Text

	performTestChecks(
		map[string]bool{
			"Blend not found by ID": updatedBlend == nil,
			fmt.Sprintf("Expected drip text to be 'updated by id', got '%s'", dripText): dripText != "updated by id",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test updating first drip in multi-drip blend
func TestUpdateDripInBlend_UpdateFirstDrip(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend with multiple drips
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Multi Blend"}, "first drip")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Multi Blend"}, "second drip")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Multi Blend"}, "third drip")

	// Update the first drip (id 0)
	tc.UpdateDripInBlend(&BlendIdentifier{Id: -1, Title: "Multi Blend"}, 0, "updated first")

	blends := tc.dataSource.Load()
	multiBlend := findBlendByTitle(blends, "Multi Blend")

	firstDripText := multiBlend.Drips[0].Text
	secondDripText := multiBlend.Drips[1].Text
	thirdDripText := multiBlend.Drips[2].Text

	performTestChecks(
		map[string]bool{
			"Multi Blend not found": multiBlend == nil,
			fmt.Sprintf("Expected first drip to be 'updated first', got '%s'", firstDripText):     firstDripText != "updated first",
			fmt.Sprintf("Expected second drip to remain 'second drip', got '%s'", secondDripText): secondDripText != "second drip",
			fmt.Sprintf("Expected third drip to remain 'third drip', got '%s'", thirdDripText):    thirdDripText != "third drip",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

// Test updating last drip in multi-drip blend
func TestUpdateDripInBlend_UpdateLastDrip(t *testing.T) {
	tc := getTestCoffeeShopController()

	// Setup: Add a blend with multiple drips
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Last Drip Blend"}, "first")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Last Drip Blend"}, "second")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Last Drip Blend"}, "third")

	// Update the last drip (id 2)
	tc.UpdateDripInBlend(&BlendIdentifier{Id: -1, Title: "Last Drip Blend"}, 2, "updated third")

	blends := tc.dataSource.Load()
	lastDripBlend := findBlendByTitle(blends, "Last Drip Blend")

	lastDripText := lastDripBlend.Drips[2].Text

	performTestChecks(
		map[string]bool{
			"Last Drip Blend not found": lastDripBlend == nil,
			fmt.Sprintf("Expected last drip to be 'updated third', got '%s'", lastDripText): lastDripText != "updated third",
		},
	)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

func getTestCoffeeShopController() *CoffeeShopController {
	tp := carafepath.GetTestingCarafePath()

	return NewCoffeeShopController(tp)
}

func performTestChecks(checks map[string]bool) {
	for check, passed := range checks {
		assert.Assert(!passed, check)
	}
}

func findBlendByTitle(blends []datamodel.Blend, title string) *datamodel.Blend {
	for i := range blends {
		if blends[i].Title == title {
			return &blends[i]
		}
	}
	return nil
}
