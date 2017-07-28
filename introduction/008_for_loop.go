package main

import "fmt"

func main() {

	intSlice := []int{1, 2, 3, 4, 5}

	//standart for loop
	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("%d: %d \n", i, intSlice[i])
	}

	//range loop
	for i, intSliceItem := range intSlice {
		fmt.Printf("%d: %d \n", i, intSliceItem)
	}

	//only indexes
	for i := range intSlice {
		fmt.Println(i)
	}

	//only values
	for _, intSliceItem := range intSlice {
		fmt.Println(intSliceItem)
	}

	//like a while

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
}
