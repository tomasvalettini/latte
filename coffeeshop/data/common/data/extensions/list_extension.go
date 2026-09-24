package extensions

import (
	"sort"
	"strconv"
)

// GetNextId returns the next available ID by finding the maximum ID
// and adding 1, starting from startId.
func GetNextId[T any](slice []T, getId func(T) int, startId int) int {
	id := startId
	for _, item := range slice {
		if getId(item) > id {
			id = getId(item)
		}
	}
	return id + 1
}

// MaxIdWidth returns the maximum string width of IDs in the slice.
func MaxIdWidth[T any](slice []T, getId func(T) int) int {
	maxWidth := 0
	for _, item := range slice {
		w := len(strconv.Itoa(getId(item)))
		if w > maxWidth {
			maxWidth = w
		}
	}
	return maxWidth
}

// FindIndexFromId returns the index of the first item with the given ID,
// or the insertion point if not found (assumes sorted by getId).
func FindIndexFromId[T any](slice []T, id int, getId func(T) int) int {
	for idx, item := range slice {
		if getId(item) == id || id < getId(item) {
			return idx
		}
	}
	return len(slice)
}

// SortById sorts the slice in ascending order by ID.
func SortById[T any](slice []T, getId func(T) int) {
	sort.Slice(slice, func(i, j int) bool {
		return getId(slice[i]) < getId(slice[j])
	})
}

// UpdateIds updates IDs starting from startIdx, incrementing consecutive IDs
// to avoid collisions.
func UpdateIds[T any](slice []T, startIdx int, idToAdd int, getId func(T) int, setId func(int, int)) {
	idToUpdate := idToAdd
	for i := startIdx; i < len(slice); i++ {
		if getId(slice[i]) == idToUpdate {
			setId(i, idToUpdate+1)
			idToUpdate++
		}
	}
}
