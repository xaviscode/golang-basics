package main

import "fmt"

//I want the pointer of this variable
func changeValue(str *string) {
	*str = "changed!"
}

//You are changing a copy of the string
func changeValue2(str string) {
	str = "changed!"
}

/*
	func changeValue2(str string) string {
		str = "changed!"
		return str
	}
*/

//& *
func main() {

	x := 7

	//Pointer
	y := &x
	fmt.Println(x, y)

	//Derefrence
	*y = 8
	fmt.Println(x, y)

	toChange := "Hello"

	fmt.Println(toChange)
	//changeValue(&toChange)
	//fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)
	//fmt.Println(changeValue2(toChange))

	var pointer *string = &toChange
	fmt.Println(pointer, &pointer, *pointer)
}
