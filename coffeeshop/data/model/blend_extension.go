package datamodel

import (
	"github.com/tomasvalettini/latte/coffeeshop/data/common/data/extensions"
)

func GetNextBlendId(blends []Blend) int {
	return extensions.GetNextId(blends, func(b Blend) int { return b.Id }, 0)
}

func MaxBlendIdWidth(blends []Blend) int {
	return extensions.MaxIdWidth(blends, func(b Blend) int { return b.Id })
}

func SortBlendsById(blends []Blend) {
	extensions.SortById(blends, func(b Blend) int { return b.Id })
}
