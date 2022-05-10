package main

import (
	"fmt"
	"time"
	//"sync"
)

func main() {

	//var wg sync.WaitGroup
	//wg.Add(1)

	//Gorutine
	//go count("sheep")
	//go count("fish")

	//time.Sleep(time.Second * 2)
	//fmt.Scanln()

	/*
		go func() {
			count("sheep")
			wg.Done()
		}()

		wg.Wait()
	*/

	c := make(chan string)
	go count("sheep", c)

	for msg := range c {
		//msg, open := <-c
		/*
			if !open {
				break
			}
		*/
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}

/*
	func count(thing string) {
		//
		//	for i := 0; true; i++ {
		//		fmt.Println(i, thing)
		//		time.Sleep(time.Millisecond * 500)
		//	}


		for i := 1; i <= 5; i++ {
			fmt.Println(i, thing)
			time.Sleep(time.Millisecond * 500)
		}
	}
*/
