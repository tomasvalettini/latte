package backlog

import (
	"log"
	"strconv"
)

func GetNextId(tasks []Task) int {
	var id int = -1

	for _, task := range tasks {
		if task.Id > id {
			id = task.Id
		}
	}

	return id + 1
}

func MaxIdWidth(tasks []Task) int {
	maxWidth := 0

	for _, task := range tasks {
		w := len(strconv.Itoa(task.Id))
		if w > maxWidth {
			maxWidth = w
		}
	}

	return maxWidth
}

func FindIndexFromId(tasks []Task, id int) int {
	for idx, task := range tasks {
		if task.Id == id {
			return idx
		}
	}

	log.Fatalln("Task id was not found")
	return -1
}
