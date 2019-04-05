package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	side float64
}

type triangle struct {
	sideOne   float64
	sideTwo   float64
	sideThree float64
}

func (t triangle) getArea() float64 {
	return t.sideOne * t.sideTwo * t.sideThree
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println("The area of this shape is:", s.getArea())
}

func main() {
	mySquare := square{side: 3}
	myTriangle := triangle{sideOne: 3, sideTwo: 4, sideThree: 5}

	printArea(mySquare)
	printArea(myTriangle)
}
