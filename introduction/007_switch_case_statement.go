package main

import "fmt"

func main()  {
   number := 1
   switch {
  case number == 1:
		fmt.Println("Selected case 1")
  case number == 2:
		fmt.Println("Selected case 2")
	default:
		fmt.Println("Whatever")
	}
}
