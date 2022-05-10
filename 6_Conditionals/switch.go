package main

import "fmt"

func main() {

	ans := 2

	fmt.Printf("%T\n", ans)

	switch ans {
	case 1, -1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	/*
		Can not print because ans is not a String
		case "Hello":
	*/
	default:
		fmt.Println("Not a case")
	}

	switch {
	case ans > 0:
		fmt.Println("Greather than 0")
	case ans < 0:
		fmt.Println("Less than 0")
	}
}
