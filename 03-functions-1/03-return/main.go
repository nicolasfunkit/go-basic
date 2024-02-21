package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func Sum(parameters ...int) int {
	total := 0
	for _, num := range parameters {
		total += num
	}

	return total
}
