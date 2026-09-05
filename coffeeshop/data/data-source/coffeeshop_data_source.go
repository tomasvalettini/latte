package datasource

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/tomasvalettini/latte/assert"
	datamodel "github.com/tomasvalettini/latte/coffeeshop/data/model"
)

type CoffeeShopDataSource struct {
	path string
}

func NewCoffeeShopDataSource(path string) *CoffeeShopDataSource {
	return &CoffeeShopDataSource{
		path: path,
	}
}

func (coffeeShop *CoffeeShopDataSource) Load() []datamodel.Blend {
	data, err := os.ReadFile(coffeeShop.path)

	if err != nil {
		if os.IsNotExist(err) {
			return []datamodel.Blend{}
		}

		log.Fatalln("Error opening coffee shop file :(.")
	}

	var blends []datamodel.Blend

	err = json.Unmarshal(data, &blends)
	assert.Assert(err == nil, "Error while parsing json.")

	return blends
}

func (coffeeShop *CoffeeShopDataSource) Save(blends []datamodel.Blend) {
	err := os.MkdirAll(filepath.Dir(coffeeShop.path), 0o755)
	assert.Assert(err == nil, "Error while creating and opening coffee shop db.")

	data, err := json.MarshalIndent(blends, "", "  ")
	assert.Assert(err == nil, "Error while creating json.")

	data = append(data, '\n')
	os.WriteFile(coffeeShop.path, data, 0o644)
}
