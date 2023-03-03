package main

import "fmt"

type Square struct {
	sideLength float64
}

type Triangle struct {
	base   float64
	height float64
}

type Shape interface {
	getArea() float64
}

func main() {
	triangle := Triangle{base: 10, height: 10}
	square := Square{sideLength: 10}

	printArea(triangle)
	printArea(square)
}

func (t Triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}
