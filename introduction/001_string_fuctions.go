package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Hello,World!"
	//To Lower Case
	lower := strings.ToLower(word)
	fmt.Println(lower)

	//To Upper Case
	upper := strings.ToUpper(word)
	fmt.Println(upper)

	//contains
	if strings.Contains(word, "World") {
		fmt.Println("Evreka!")
	}

	// substring
	fmt.Println("Only printing out characters 0 to 4 : " + word[0:4])

	fmt.Println("Only first three Ccharacters: " + word[:3])

	// split string
	fullSentence := "any fool can write code that a computer can understand"
	splittedWords := strings.Split(fullSentence, " ")
	fmt.Printf("%v \n", splittedWords)

}
