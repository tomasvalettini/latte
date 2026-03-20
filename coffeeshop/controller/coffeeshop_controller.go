package controller

import (
	"fmt"
	"strings"

	datasource "github.com/tomasvalettini/latte/coffeeshop/data/data-source"
	carafepath "github.com/tomasvalettini/latte/coffeeshop/data/data-source/path"
	datamodel "github.com/tomasvalettini/latte/coffeeshop/data/model"
)

// default blend (collection) that is `latte` flavoured
const house_blend_title = "House Blend"
const house_blend_id = 0

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

	if bi == nil {
		// show all blends
		for _, blend := range blends {
			printBlend(blend)
		}
		return
	}

	if !bi.IsValid() {
		fmt.Println("Blend id or Blend title should be specified")
		return
	}

	printBlend(*getBlendFromIdentifier(blends, bi))
}

func (csc *CoffeeShopController) AddToBlends(bi *BlendIdentifier, dripText string) {
	if dripText == "" {
		fmt.Println("Missing text for drip. Please specify the text to add the drip.")
		return
	}

	blends := csc.dataSource.Load()
	var foundBlend *datamodel.Blend

	if bi == nil {
		foundBlend = getOrCreateBlendFromIdentifier(
			blends,
			&BlendIdentifier{
				Id:    house_blend_id,
				Title: house_blend_title,
			},
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

	if bi.IsTitleValid() {
		return &datamodel.Blend{
			Id:    datamodel.GetNextBlendId(blends),
			Title: bi.Title,
		}
	}

	// adding to default blend
	return &datamodel.Blend{
		Id:    house_blend_id,
		Title: house_blend_title,
		Drips: []datamodel.Drip{},
	}
}

func getBlendFromIdentifier(blends []datamodel.Blend, bi *BlendIdentifier) *datamodel.Blend {
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

func printBlend(blend datamodel.Blend) {
	width := len(blend.Title) + (margin_width * 2)
	bannerLine := strings.Repeat("=", width)
	offset := strings.Repeat(" ", margin_width)
	idWidth := datamodel.MaxIdWidth(blend.Drips)

	fmt.Println(bannerLine)
	fmt.Printf("%s%s%s\n", offset, blend.Title, offset)
	fmt.Println(bannerLine)

	margin := "  "
	for _, d := range blend.Drips {
		fmt.Printf("%s[%*d] %s\n", margin, idWidth, d.Id, d.Text)
	}

	fmt.Println()
}
