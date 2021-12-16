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

		go func(i int) {

			wg.Add(1)

			fmt.Println(i)

			wg.Done()

		}(i)

		if i == 9 {
			over <- true
		}

	}

	<-over

	wg.Wait()

	fmt.Println("Over")

}
