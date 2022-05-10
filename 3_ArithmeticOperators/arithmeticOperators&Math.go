package main

import (
	"fmt"
	"math"
)

/*
	Types:
	+ - / * % ()
*/
func main() {

	//var num1 float64 = 8
	var num1 int = 9
	var num2 int = 4
	//They need to be the same type and size
	//answer := num1 / float64(num2)
	answer := num1 / num2
	fmt.Printf("%d\n", answer)

	var num3 float32 = 9
	var num4 float32 = 4
	answer1 := num3 / num4
	fmt.Printf("%g\n", answer1)

	var sqrt float64 = math.Sqrt(9)
	fmt.Println("The sqrt of 9 is:", sqrt)
}
