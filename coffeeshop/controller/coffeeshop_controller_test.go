package controller

import (
	"fmt"
	"os"
	"testing"

	carafepath "github.com/tomasvalettini/latte/coffeeshop/data/data-source/path"
)

func TestCoffeeShopController(t *testing.T) {
	tc := getTestCoffeeShopController()

	// listing all blends
	tc.ListBlends(nil)
	printSeparator()

	tc.AddToBlends(nil, "")
	tc.AddToBlends(nil, "test drip")
	tc.AddToBlends(nil, "test drip")
	tc.AddToBlends(nil, "test drip")

	tc.ListBlends(nil)
	printSeparator()
	tc.ListBlends(&BlendIdentifier{Id: 0, Title: "House Blend"})
	printSeparator()

	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "testing adding espresso")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "testing adding espresso")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "testing adding espresso")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Espresso"}, "testing adding espresso")

	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Capuccino"}, "testing adding capuccino")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Capuccino"}, "testing adding capuccino")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Capuccino"}, "testing adding capuccino")
	tc.AddToBlends(&BlendIdentifier{Id: -1, Title: "Capuccino"}, "testing adding capuccino")

	tc.ListBlends(nil)

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})
}

func getTestCoffeeShopController() *CoffeeShopController {
	tp := carafepath.GetTestingCarafePath()

	return NewCoffeeShopController(tp)
}

func printSeparator() {
	fmt.Println("---------------------------------------")
	fmt.Println()
}
