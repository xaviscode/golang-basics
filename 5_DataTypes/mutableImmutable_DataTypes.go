package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 100
}

func main() {
	/*
		//Immutable data type - Default
		var x int = 5
		y := x
		y = 7
		fmt.Println(x, y)
	*/

	/*
		//Mutable data type - Slices
		var x []int = []int{3, 4, 5}
		y := x
		y[0] = 100
		fmt.Println(x, y)
	*/
	/*
		//Mutable data type - map
		var x map[string]int = map[string]int{"Hello": 3}
		y := x
		y["y"] = 100
		x["x"] = 50
		fmt.Println(x, y)
	*/

	/*
		var x [2]int = [2]int{3, 4}
		y := x
		y[0] = 100
		fmt.Println(x, y)
	*/

	var x []int = []int{3, 4, 5}
	fmt.Println(x)
	changeFirst(x)
	fmt.Println(x)
}
