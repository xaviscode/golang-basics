package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born: ")
	scanner.Scan()
	/*
		input := scanner.Text()
		cannot use 2020 (type untyped int) as type string
		fmt.Printf("You will be %d year old at the end of 2020", 2020-input)
	*/

	//This underscore is like an error, when you do not return an Int value the input variable will be empty.
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d year old at the end of 2020 \n", 2020-input)
}
