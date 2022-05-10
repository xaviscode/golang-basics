package main

import "fmt"

/*
	< > <= >= == !=
*/
func main() {

	x := 5
	y := 6.5
	/*
		x := 5
		y := 6.5
		invalid operation: x == y (mismatched types int and float64)
		val := x == y
	*/
	val := float64(x)+1.5 == y
	fmt.Printf("%t\n", val)

	//Comparing strings
	a := "Hello"
	b := "Hello"
	c := "Hola"
	d := "A"
	e := "a"

	valS := a == b
	fmt.Printf("%t %t %t\n", valS, a == c, d < e)

	//cannot convert e (type string) to type int
	//fmt.Println(int(e))
}
