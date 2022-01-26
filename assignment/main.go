package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base float64
}

func main() {
	tri := triangle{
		height: 4,
		base: 2,
	}
	sqr := square{
		sideLength: 5,
	}

	printArea(tri)
	printArea(sqr)
}

func printArea(s shape) {
	fmt.Println("The area for the shape is",  s.getArea())
}

func (t triangle) getArea() (float64) {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}