package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	//pow(x int) float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{4.5}
	r1 := rect{5, 7}

	/*
		We can not do that
		shapes := []type{c1, r1}
		shapes[1].area()
	*/

	shapes := []shape{c1, r1}

	//fmt.Println(c1.area())

	for _, shape := range shapes {
		//NO: fmt.Println(shape.radius)
		fmt.Println(shape.area())
	}

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
