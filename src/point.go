package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}


func (p *Point) moveTo(x int, y int) {
	p.X += x
	p.Y += y 
}


func main() {
	p := Point{2, 4}

	p.moveTo(5, 7)

	fmt.Println(p.X, p.Y)
}      