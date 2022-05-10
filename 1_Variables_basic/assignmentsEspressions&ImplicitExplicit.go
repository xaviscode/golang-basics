package main

import "fmt"

func main() {

	/*
		Explicit define variable is when
		you define the type of the variable.
		var number uint16 = 260
	*/

	// Implicit define variable is when you do not define the type.
	var number = 200.1

	// It shows the type of the variable.
	fmt.Printf("%T", number)

	// Assingn a variable.
	number1 := 6
	fmt.Printf("%T", number1)

	// Booleans have to assign as 'false' not 'False'
	// boolean := false

	// When you not define the value of the variable it gives you a 0.
	var num int16
	var bl bool
	fmt.Println(num, bl)
}
