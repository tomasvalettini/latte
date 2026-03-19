package dripdatasource

import dripdatamodel "github.com/tomasvalettini/latte/drips/data/model"

type DripDataSource interface {
	Load() []dripdatamodel.Drip
	Save(drips []dripdatamodel.Drip)
}
