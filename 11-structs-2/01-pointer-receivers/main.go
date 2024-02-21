package main

type Position struct {
	X int
	Y int
}

func (p *Position) Move(x int, y int) {
	p.X = p.X + x
	p.Y = p.Y + y
}

func main() {
	p := Position{X: 10, Y: 20}
	p.Move(5, -5)
}
