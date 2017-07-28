package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	isExists := IsExists(1, intSlice)

	fmt.Println(isExists)
}

func IsExists(item int, items []int) bool {
	for _, currentItem := range items {
		if item == currentItem {
			return true
		}
	}
	return false
}
