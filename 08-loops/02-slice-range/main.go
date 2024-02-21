package main

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

func Sum(values ...int) int {
	total := 0
	for _, value := range values {
		total = total + value
	}
	return total
}
