package main

import "fmt"

func main() {

	var a []int = []int{1, 6, 3, 4, 5, 6, 7, 8, 3}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	//element == a[i]
	for i, element := range a {
		fmt.Printf("%d: %d\n", i, element)
	}

	/*
		If i do not want to use a variable we can just _ like this:
		for _, element := range a {
			fmt.Printf("%d\n", element)
		}
		We are not using i then we just change it to _
	*/

	//Example1
	for i, y := range a {
		for j, y2 := range a {
			if y == y2 && i > j {
				fmt.Println(y)
			}
		}
	}

	//Example2
	for i, y := range a {
		for j := i + 1; j < len(a); j++ {
			y2 := a[j]
			if y2 == y {
				fmt.Println(y)
			}
		}
	}
}
