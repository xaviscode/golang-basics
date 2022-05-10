package main

import "fmt"

func test2(x int) {
	fmt.Println("Hello", x)
}

func test4(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {

	//We can assign functions in to a value
	x := test2
	x(5)

	//Another way to do it
	test3 := func() {
		fmt.Println("Test3")
	}
	test3()

	multNeg := func(x int) int {
		return x * -1
	}
	/*
		multNeg := func(x int) int {
			return x * -1
		}(8)
		fmt.Println(multNeg)
	*/
	test4(multNeg)

	test5 := func(x int) int {
		return x * 7
	}
	test4(test5)

	returnFunc("Hi")()
	y := returnFunc("Hello")
	y()
}
