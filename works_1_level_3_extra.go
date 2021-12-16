package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var num []int

	num = make([]int, 0)

	for k := 0; k < 100; k++ {
		num = append(num, rand.Intn(1000))
	}

	rank(num, 100)

	for k := 0; k < 100; k++ {
		fmt.Println(k, num[k])
	}

}

func rank(num []int, i int) {

	for k := 0; k < i; k++ {

		for q := 0; q < i-1; q++ {

			if num[q] < num[q+1] {

				var temp = num[q]

				num[q] = num[q+1]

				num[q+1] = temp

			}

		}

	}

}
