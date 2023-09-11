package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Cicle struct {
	radius float64
}

func (c Cicle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func PrintArea(s Shape) {
	fmt.Printf("Area:%0.2f", s.Area())
}
func main() {
	var circlerad float64
	fmt.Print("enter the radius:\n")
	fmt.Scan(&circlerad)
	circle := Cicle{radius: circlerad}
	PrintArea(circle)
	var rectwidth, recheight float64
	fmt.Print("\nenter the width:\n")
	fmt.Scan(&rectwidth)
	fmt.Print("\nenter the height:\n")
	fmt.Scan(&recheight)
	rectangle := Rectangle{width: rectwidth, height: recheight}
	PrintArea(rectangle)

}
