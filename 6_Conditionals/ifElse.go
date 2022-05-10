package main

import "fmt"

func main() {

	name := "Helo"

	fmt.Println("Before if")
	if name == "Hello" {
		fmt.Println("Inside if")
	} else {
		fmt.Println("Inside else")
	}
	fmt.Println("After if")

	/*
		if condition {

		} else if condition {

		} else if condition {

		} else {

		}
	*/
}
