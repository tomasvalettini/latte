package dripdatamodel

import (
	"log"
	"strconv"
)

func GetNextId(drips []Drip) int {
	var id int = -1

	for _, drip := range drips {
		if drip.Id > id {
			id = drip.Id
		}
	}

	return id + 1
}

func MaxIdWidth(drips []Drip) int {
	maxWidth := 0

	for _, drip := range drips {
		w := len(strconv.Itoa(drip.Id))
		if w > maxWidth {
			maxWidth = w
		}
	}

	return maxWidth
}

func FindIndexFromId(drips []Drip, id int) int {
	for idx, drip := range drips {
		if drip.Id == id {
			return idx
		}
	}

	log.Fatalln("Drip id was not found")
	return -1
}
