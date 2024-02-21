package main

import "fmt"

var cpt = 0

func AllocateBuffer() *string {
	if cpt >= 3 {
		return nil
	}
	cpt++
	return new(string)
}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}
