package main

import "fmt"

func main() {

	//No different types and just indicate the Max size
	var arrI [5]int
	fmt.Println(arrI)
	var arrS [10]string
	fmt.Println(arrS)

	/*
		var arr [5]int
		[0 0 0 0 0]
		 0 1 2 3 4
	*/

	var arr [5]int
	arr[0] = 100
	arr[4] = 400
	fmt.Println(arr)

	//It is required a size
	arr1 := [3]int{4, 5, 6}
	fmt.Println(arr1)
	fmt.Println(len(arr1))

	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	fmt.Println(sum)

	//Arrays inside Arrays
	//{{}, {}, {}}
	arr2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2D[0][2])
}
