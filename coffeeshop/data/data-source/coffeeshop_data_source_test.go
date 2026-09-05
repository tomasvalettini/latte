package datasource_test

import (
	"os"
	"testing"

	"github.com/tomasvalettini/latte/assert"
	datasource "github.com/tomasvalettini/latte/coffeeshop/data/data-source"
	carafepath "github.com/tomasvalettini/latte/coffeeshop/data/data-source/path"
	datamodel "github.com/tomasvalettini/latte/coffeeshop/data/model"
	testutils "github.com/tomasvalettini/latte/test-utils"
)

func TestCoffeeShopDataSourceLogic(t *testing.T) {
	cPath := carafepath.GetTestingCarafePath()
	coffeeShopDataSource := datasource.NewCoffeeShopDataSource(cPath.GetCarafePath())
	blends := coffeeShopDataSource.Load()

	assert.Assert(len(blends) == 0, "There should not be any blends yet!")

	t.Cleanup(func() {
		os.RemoveAll(carafepath.TMP)
	})

	testBlends := []datamodel.Blend{
		{
			Id:    1,
			Title: "Blend 1",
			Drips: []datamodel.Drip{
				{Id: 1, Text: "test"},
				{Id: 2, Text: "test"},
				{Id: 3, Text: "test"},
			},
		},
		{
			Id:    2,
			Title: "Blend 2",
			Drips: []datamodel.Drip{
				{Id: 1, Text: "test"},
				{Id: 2, Text: "test"},
				{Id: 3, Text: "test"},
			},
		}, {
			Id:    4,
			Title: "Empty Blend",
			Drips: []datamodel.Drip{}, // Good for testing edge cases where a blend has no drips
		},
	}

	coffeeShopDataSource.Save(testBlends)
	blends = coffeeShopDataSource.Load()
	assert.Assert(len(blends) != 0, "There should be blends saved!")
}

func TestCoffeeShopDataSourceFailingFile(t *testing.T) {
	testutils.RequireExit(t, "TestCoffeeShopDataSourceFailingFile", testingFailingFile)
}

func testingFailingFile() {
	coffeeShopDataSource := datasource.NewCoffeeShopDataSource(carafepath.TMP)
	coffeeShopDataSource.Load()
}
