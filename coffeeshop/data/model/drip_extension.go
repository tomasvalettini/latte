package datamodel

import (
	"github.com/tomasvalettini/latte/coffeeshop/data/common/data/extensions"
)

func GetNextId(drips []Drip) int {
	return extensions.GetNextId(drips, func(d Drip) int { return d.Id }, -1)
}

func MaxDripIdWidth(drips []Drip) int {
	return extensions.MaxIdWidth(drips, func(d Drip) int { return d.Id })
}

// this method assumes the drip array is sorted
func FindIndexFromId(drips []Drip, id int) int {
	return extensions.FindIndexFromId(drips, id, func(d Drip) int { return d.Id })
}

func UpdateDripsIds(drips []Drip, idx int, idToAdd int) {
	extensions.UpdateIds(drips, idx, idToAdd,
		func(d Drip) int { return d.Id },
		func(i int, id int) { drips[i].Id = id })
}

func SortDripsById(drips []Drip) {
	extensions.SortById(drips, func(d Drip) int { return d.Id })
}
