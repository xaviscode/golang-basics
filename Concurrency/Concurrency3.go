package main

import (
	"fmt"
	"time"
)

//It execute every 2 seconds because the c2 blocks the c1
func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}
