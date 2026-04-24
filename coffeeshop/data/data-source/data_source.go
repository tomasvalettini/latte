package datasource

import datamodel "github.com/tomasvalettini/latte/coffeeshop/data/model"

type DataSource interface {
	Load() []datamodel.Blend
	Save(blends []datamodel.Blend)
}
