package main

import "fmt"

func main() {

	/*
		These format are named 'verbs' and are
		derived from C but simpler.
	*/

	// Printing the values
	fmt.Printf("Hello %T %v \n", 10, 10)

	/*
		You can save a print in to a variable
		var x string = fmt.Sprintf(..., ...)
	*/

	fmt.Printf("General: %v \n", "Hello")
	fmt.Printf("Boolean: %t \n", false)
	fmt.Printf("Integer: %b %x %X %d %o %O %U \n", 2, 45, 45, 45, 45, 45, 45)
	fmt.Printf("Floating: %e %E %f %g %x %X \n", 45.243434, 45.243434, 45.243434, 45.243434, 45.2434463463757254783547364783, 45.243434)
	fmt.Printf("String: %s %q %x %X \n", "Hello", "Hello", "Hello", "Hello")
	fmt.Printf("Precision: %19b, %9.5f, %v \n", 2, 5.555, 56)

	for ch := 75; ch <= 100; ch++ {
		fmt.Printf("ASCII value = %d, Character = %c\n", ch, ch)
	}
}
