package main

import (
	"fmt"
)

type Point struct {
	x        int32
	y        int32
	isOnGrid bool
}

type Point1 struct {
	x int32
	y int32
}

type Circle struct {
	radius float64

	//You can give a name or not
	center *Point1
	//*Point1
}

//If you do not pass the pointer, the function make a copy of it
func changeX(pt *Point) {
	pt.x = 100
}

func main() {
	//var p1 Point = Point{1, 2, true}
	//var p2 Point = Point{-5, 7, false}
	p1 := Point{1, 2, true}
	p2 := &Point{x: 4}

	fmt.Println(p1, p2)
	changeX(p2)

	fmt.Println(p1, p2)

	// You do not need to put a Derefrence
	p3 := Point{y: 3}
	p3.x = 8
	fmt.Println(p3)

	//p11 := &Point1{y: 3}
	//c1 := Circle{4.56, p11}

	c2 := Circle{4.56, &Point1{4, 5}}

	//Accesing x,y from c2
	fmt.Println(c2)
	fmt.Println(c2.radius)
	fmt.Println(c2.center)
	fmt.Println(c2.center.x)
	//fmt.Println(c2.x)
}
