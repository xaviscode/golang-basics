package main

import "fmt"

func main() {

	var x [5]int = [5]int{1, 2, 3, 4, 5}
	//When there is no values inside the square brackets, it is a slice.
	var s []int = x[1:3]

	fmt.Println(s, len(s), cap(s), s[:cap(s)])

	var a []int = []int{5, 6, 7, 8, 9}
	fmt.Println(cap(a[:3]))

	//Add variables to a slice
	b := append(a, 10)
	a = append(a, 10)
	fmt.Println(b)
	fmt.Println(a)

	y := make([]int, 5)
	fmt.Println(y)
	//If it were an array it will shows as [5]int
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", x)
}
