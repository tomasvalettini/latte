package backlog

import (
	"log"
	"strconv"
)

func GetNextId(items []Item) int {
	var id int = -1

	for _, item := range items {
		if item.Id > id {
			id = item.Id
		}
	}

	return id + 1
}

func MaxIdWidth(items []Item) int {
	maxWidth := 0

	for _, item := range items {
		w := len(strconv.Itoa(item.Id))
		if w > maxWidth {
			maxWidth = w
		}
	}

	return maxWidth
}

func FindIndexFromId(items []Item, id int) int {
	for idx, item := range items {
		if item.Id == id {
			return idx
		}
	}

	log.Fatalln("Item id was not found")
	return -1
}
