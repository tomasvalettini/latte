package datamodel

func GetNextBlendId(blends []Blend) int {
	var id int = -1

	for _, blend := range blends {
		if blend.Id > id {
			id = blend.Id
		}
	}

	return id + 1
}
