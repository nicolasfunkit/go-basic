package main

import (
	"fmt"
	"os"
)

func CreateDirectory(name string) bool {
	error := os.Mkdir(name, 0755)
	return error == nil

}

func main() {
	ok := CreateDirectory("my-directory")
	if ok {
		fmt.Println("Directory created")
	} else {
		fmt.Println("Failed to create directory")
	}
}
