package main

import "fmt"

func main() {
	sum := Sum(1, 2)
	fmt.Println(sum)
}

func Sum(number1, number2 int32) int32 {
	return number1 + number2
}
