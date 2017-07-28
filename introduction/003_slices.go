package main

import "fmt"

func main() {

	//init with sample data
	stringArray := []string{"cem", "ilkay", "hakan", "emine"}
	fmt.Printf("%v \n", stringArray)

	//as a standart decleration
	var intSlice []int

	//add value to current slice
	intSlice = append(intSlice, 1)
	intSlice = append(intSlice, 2)

	//add multiple values
	intSlice = append(intSlice, 3, 4, 5)
	fmt.Printf("%v \n", intSlice)

	// get length of a slice
	fmt.Println("Length: ", len(intSlice))

	// retrieve element by index
	fmt.Println(intSlice[3])

	// create a new slice from current slice
	newIntSlice := intSlice[1:3]
	fmt.Printf("%v \n", newIntSlice)
}
