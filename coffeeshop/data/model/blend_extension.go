package datamodel

func GetNextBlendId(blends []Blend) int {
	var id int = 0 // 0 is for the default blend

	for _, blend := range blends {
		if blend.Id > id {
			id = blend.Id
		}
	}

	return id + 1
}
