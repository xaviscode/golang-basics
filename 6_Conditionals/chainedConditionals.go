package main

import "fmt"

//It combines multiple boolean expresions (AND-&& OR-|| NOT-!)
func main() {

	x := 8
	val := (true || false) && false || (x < 9)
	fmt.Printf("%t\n", val)
}
