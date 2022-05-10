package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	x := 1
	//There is no while loops
	for x < 3 {
		fmt.Println("I am a while loop")
		//x+=1
		x++
	}

	for {
		//Infinite loop
		break    //It just break the loop
		continue //It immediately jumps to the top of the loop
	}

	fmt.Println("---------------------------------------")
	for i := 0; i <= 100; i++ {
		if i != 0 && i%3 == 0 && i%7 == 0 && i%9 == 0 {
			fmt.Println(i)
			continue
		}
		fmt.Println("N")
	}
	fmt.Println("---------------------------------------")
	//OR
	for i := 0; i <= 100; i++ {
		if i != 0 && i%3 == 0 && i%7 == 0 && i%9 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println("N")
		}
	}
	fmt.Println("---------------------------------------")
	//Print the first value
	for i := 0; i <= 100; i++ {
		if i != 0 && i%3 == 0 && i%7 == 0 && i%9 == 0 {
			fmt.Println(i)
			break
		}
	}
}
