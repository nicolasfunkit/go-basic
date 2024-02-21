package main

import (
	"fmt"
	"os"
)

func main() {
	ok := CheckFile("input.csv")
	if ok {
		fmt.Println("File correctly read")
	} else {
		fmt.Println("Failed to read file")
	}
}

func CheckFile(filename string) bool {
	_, error := os.ReadFile(filename)
	return error == nil
}
