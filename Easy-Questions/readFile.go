package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileContent, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		fmt.Print("Error while reading the file : ", err)
		return
	}
	fmt.Print(string(fileContent))
}
