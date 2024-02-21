package main

import (
	"os"
)

func main() {
	if _, err := os.Stat("analysis.xlsx"); err != nil {
		fmt.println(err.Error())
		return
	}

	fmt.println("ok")
}
