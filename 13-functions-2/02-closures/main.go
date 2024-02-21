package main

import "fmt"

func WordGenerator(words []string) func() string {
	counter := 0
	return func() string {
		if counter >= len(words) {
			counter = 0
		}

		word := words[counter]
		counter++
		return word
	}
}

func main() {
	continents := []string{
		"Africa",
		"Antarctica",
		"Asia",
		"Australia",
		"Europe",
		"North America",
		"South America",
	}

	generator := WordGenerator(continents)

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
