package main

import "fmt"

func main() {

	//Map do not store in order the elements
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	//mpp := make(map[string]int)
	fmt.Println(mp)
	fmt.Println(mp["apple"])

	//Change values
	mp["apple"] = 900
	fmt.Println(mp)

	//Add values
	mp["Hello"] = 900
	fmt.Println(mp)

	//Delete values
	delete(mp, "apple")
	fmt.Println(mp)

	//Check if there is the key
	val, ok := mp["pear"]
	fmt.Println(val, ok)

	//We can use the len() in maps
	fmt.Println(len(mp))

}
