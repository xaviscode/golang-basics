package main

import "fmt"

//Example: x -> |x|

func test() {
	fmt.Println("Test")
}

func test1(x int) {
	fmt.Println(x)
}

//You can use only the name of the value if all there are the same type
func add(x, y int) int {
	return (x + y)
}

//You can return more than 1 value this way:
func addSubs(x, y int) (int, int) {
	return x + y, x - y
}

func multDiv(x, y float32) (z1, z2 float32) {
	//It defer the value until the function is done
	defer fmt.Println("Hello")
	z1 = x * y
	z2 = x / y
	fmt.Println("Before return")
	return
}

//We can call a function as many times we want
func main() {
	test()
	test()
	test()
	test1(5)
	ans := add(5, 9)
	fmt.Println(ans)

	ans1, ans2 := addSubs(10, 5)
	fmt.Println(ans1, ans2)

	ans3, ans4 := multDiv(2, 2)
	fmt.Println(ans3, ans4)
}
