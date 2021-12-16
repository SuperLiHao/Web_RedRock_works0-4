package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var mu sync.Mutex

var i int

func main() {

	over := make(chan bool, 1)

	for i = 0; i < 10; i++ {

		wg.Add(1)

		go func() {

			fmt.Println(i)

			wg.Done()

		}()

		wg.Wait()

		if i == 9 {
			over <- true
		}

	}
	<-over
	fmt.Println("Over")

}
