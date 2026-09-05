package controller

import (
	"fmt"
	"strconv"
	"strings"

	datasource "github.com/tomasvalettini/latte/coffeeshop/data/data-source"
	carafepath "github.com/tomasvalettini/latte/coffeeshop/data/data-source/path"
	datamodel "github.com/tomasvalettini/latte/coffeeshop/data/model"
)

// default blend (collection) that is `latte` flavoured
const HOUSE_BLEND_TITLE = "House Blend"
const HOUSE_BLEND_ID = 0

// constants used in showing the blends and drips
const margin_width = 6

type CoffeeShopController struct {
	dataSource *datasource.CoffeeShopDataSource
}

func NewCoffeeShopController(path carafepath.CarafePath) *CoffeeShopController {
	ds := datasource.NewCoffeeShopDataSource(path.GetCarafePath())

	return &CoffeeShopController{
		dataSource: ds,
	}
}

func (csc *CoffeeShopController) ListBlends(bi *BlendIdentifier) {
	blends := csc.dataSource.Load()

	if len(blends) == 0 {
		fmt.Println("Nothing to show yet!")
		return
	}

	if bi != nil {
		bi = bi.Validate()
	}

	if bi == nil {
		// show all blends summary
		printBlends(blends)
		return
	}

	if !bi.IsValid() {
		fmt.Println("Blend id or Blend title should be specified")
		return
	}

	printBlendDrips(*getBlendFromIdentifier(blends, bi))
}

func (csc *CoffeeShopController) AddToBlends(bi *BlendIdentifier, dripText string) {
	if dripText == "" {
		fmt.Println("Missing text for drip. Please specify the text to add the drip.")
		return
	}

	blends := csc.dataSource.Load()
	var foundBlend *datamodel.Blend

	if bi != nil {
		bi = bi.Validate()
	}

	if bi == nil {
		foundBlend = getOrCreateBlendFromIdentifier(
			blends,
			nil,
		)
	} else {
		foundBlend = getOrCreateBlendFromIdentifier(blends, bi)
	}

	newDrip := datamodel.Drip{
		Id:   datamodel.GetNextId(foundBlend.Drips),
		Text: dripText,
	}

	foundBlend.Drips = append(foundBlend.Drips, newDrip)
	blends = addBlendToBlendList(blends, foundBlend)

	csc.dataSource.Save(blends)
}

func (csc *CoffeeShopController) DeleteFromBlends(bi *BlendIdentifier, dripId int) {
	if bi != nil {
		bi = bi.Validate()
	}
	// if bi is nil, default to house blend
	if bi == nil {
		bi = &BlendIdentifier{
			Id:    HOUSE_BLEND_ID,
			Title: HOUSE_BLEND_TITLE,
		}
	}

	blends := csc.dataSource.Load()
	foundBlend := getBlendFromIdentifier(blends, bi)

	if foundBlend == nil {
		fmt.Println("Blend not found.")
		return
	}

	// Requirement 2 & 3: Check if dripId is valid
	if dripId < 0 {
		// dripId is invalid, delete whole blend with confirmation
		fmt.Printf("Are you sure you want to delete the entire '%s' blend? (yes/no): ", foundBlend.Title)
		var confirmation string
		fmt.Scanln(&confirmation)

		if strings.ToLower(confirmation) == "yes" {
			// Remove the blend from the list
			var updatedBlends []datamodel.Blend
			for _, blend := range blends {
				if blend.Title != foundBlend.Title {
					updatedBlends = append(updatedBlends, blend)
				}
			}
			csc.dataSource.Save(updatedBlends)
			fmt.Printf("Blend '%s' deleted successfully.\n", foundBlend.Title)
		} else {
			fmt.Println("Deletion cancelled.")
		}
		return
	}

	// dripId is valid, delete specific drip
	dripIndex := -1
	for i, drip := range foundBlend.Drips {
		if drip.Id == dripId {
			dripIndex = i
			break
		}
	}

	if dripIndex == -1 {
		fmt.Println("Drip with id", dripId, "not found in the blend.")
		return
	}

	// Remove the drip from the slice
	foundBlend.Drips = append(foundBlend.Drips[:dripIndex], foundBlend.Drips[dripIndex+1:]...)

	// Update the blends list
	blends = addBlendToBlendList(blends, foundBlend)

	csc.dataSource.Save(blends)
	fmt.Printf("Drip with id %d deleted successfully.\n", dripId)
}

func (csc *CoffeeShopController) UpdateDripInBlend(bi *BlendIdentifier, dripId int, dripText string) {
	if bi != nil {
		bi = bi.Validate()
	}

	if bi == nil {
		fmt.Println("Blend identifier is required for updating a drip.")
		return
	} else {
		if dripId < 0 && dripText == "" {
			// modify blend title
			blends := csc.dataSource.Load()
			var foundBlend *datamodel.Blend = nil
			var blendIndex int = -1
			for i, b := range blends {
				if b.Id == bi.Id {
					foundBlend = &b
					blendIndex = i
					break
				}
			}

			if foundBlend == nil && blendIndex == -1 {
				fmt.Println("Blend not found.")
				return
			}

			foundBlend.Title = bi.Title
			blends[blendIndex].Title = bi.Title
			csc.dataSource.Save(blends)
			fmt.Printf("Blend title updated successfully to '%s'.\n", bi.Title)
			return
		}
	}

	if dripId < 0 {
		fmt.Println("Invalid drip ID. Please specify a valid drip ID to update.")
		return
	}

	if dripText == "" {
		fmt.Println("Missing text for drip. Please specify the text to update the drip.")
		return
	}

	blends := csc.dataSource.Load()
	foundBlend := getBlendFromIdentifier(blends, bi)

	if foundBlend == nil {
		fmt.Println("Blend not found.")
		return
	}

	// dripId is valid, update specific drip
	dripIndex := -1
	for i, drip := range foundBlend.Drips {
		if drip.Id == dripId {
			dripIndex = i
			break
		}
	}

	if dripIndex == -1 {
		fmt.Println("Drip with id", dripId, "not found in the blend.")
		return
	}

	foundBlend.Drips[dripIndex].Text = dripText

	// Update the blends list
	blends = addBlendToBlendList(blends, foundBlend)
	csc.dataSource.Save(blends)
}

func addBlendToBlendList(blends []datamodel.Blend, foundBlend *datamodel.Blend) []datamodel.Blend {
	index := -1

	for i, blend := range blends {
		if blend.Title == foundBlend.Title {
			index = i
			break
		}
	}

	if index == -1 {
		blends = append(blends, *foundBlend)
	} else {
		blends[index] = *foundBlend
	}

	return blends
}

func getOrCreateBlendFromIdentifier(blends []datamodel.Blend, bi *BlendIdentifier) *datamodel.Blend {
	foundBlend := getBlendFromIdentifier(blends, bi)
	if foundBlend != nil {
		return foundBlend
	}

	if bi != nil && bi.IsTitleValid() {
		return &datamodel.Blend{
			Id:    datamodel.GetNextBlendId(blends),
			Title: bi.Title,
		}
	}

	// adding to default blend - try to find existing house blend first
	for i := range blends {
		if blends[i].Id == HOUSE_BLEND_ID || blends[i].Title == HOUSE_BLEND_TITLE {
			return &blends[i]
		}
	}

	// if not found, create a new one
	return &datamodel.Blend{
		Id:    HOUSE_BLEND_ID,
		Title: HOUSE_BLEND_TITLE,
		Drips: []datamodel.Drip{},
	}
}

func getBlendFromIdentifier(blends []datamodel.Blend, bi *BlendIdentifier) *datamodel.Blend {
	if bi == nil {
		return nil
	}

	if bi.IsIdValid() && bi.IsTitleValid() {
		for _, b := range blends {
			if b.Id == bi.Id && b.Title == bi.Title {
				return &b
			}
		}
	} else if bi.IsIdValid() {
		for _, b := range blends {
			if b.Id == bi.Id {
				return &b
			}
		}
	} else if bi.IsTitleValid() {
		for _, b := range blends {
			if b.Title == bi.Title {
				return &b
			}
		}
	}

	return nil
}

func printBlends(blends []datamodel.Blend) {
	maxWidth := 15

	type blendLine struct {
		id    int
		title string
		count int
		noun  string
		full  string
	}

	lines := make([]blendLine, len(blends))

	for i, b := range blends {
		noun := "drips"
		if len(b.Drips) == 1 {
			noun = "drip"
		}

		idWidth := datamodel.MaxIdWidth(b.Drips)
		formatted := fmt.Sprintf("  [%*d] %s (%d %s)", idWidth, b.Id, b.Title, len(b.Drips), noun)

		if len(formatted) > maxWidth {
			maxWidth = len(formatted)
		}

		lines[i] = blendLine{b.Id, b.Title, len(b.Drips), noun, formatted}
	}

	title := "ALL BLENDS"
	bannerLine := strings.Repeat("=", maxWidth)

	if len(title)+4 > maxWidth {
		maxWidth = len(title) + 4
		bannerLine = strings.Repeat("=", maxWidth)
	}

	leftPad := (maxWidth - len(title)) / 2
	rightPad := maxWidth - len(title) - leftPad

	fmt.Println(bannerLine)
	fmt.Printf("%s%s%s\n", strings.Repeat(" ", leftPad), title, strings.Repeat(" ", rightPad))
	fmt.Println(bannerLine)

	for _, l := range lines {
		fmt.Println(l.full)
	}
}

func printBlendDrips(blend datamodel.Blend) {
	blendIdWidth := len(strconv.Itoa(blend.Id))
	width := len(blend.Title) + blendIdWidth + 3 + (margin_width * 2)
	bannerLine := strings.Repeat("=", width)
	offset := strings.Repeat(" ", margin_width)
	idWidth := datamodel.MaxIdWidth(blend.Drips)

	fmt.Println(bannerLine)
	fmt.Printf("%s%s (%d)%s\n", offset, blend.Title, blend.Id, offset)
	fmt.Println(bannerLine)

	margin := "  "
	for _, d := range blend.Drips {
		fmt.Printf("%s[%*d] %s\n", margin, idWidth, d.Id, d.Text)
	}

	fmt.Println()
}
