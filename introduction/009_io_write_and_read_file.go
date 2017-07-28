package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	filePath := "./data/math.txt"
  //read file
	//if exception occurred show error
	txtContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln("An exception occurred while reading file", filePath)
	}

  fmt.Println(string(txtContent))


 //write to file
  outfile := "sample.txt"
	err = ioutil.WriteFile(outfile, txtContent, 0644)
	if err != nil {
		log.Fatalln("An exception occurred while writing file: ", err)
	}
}
