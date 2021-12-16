package main

import (
	"fmt"
	"sort"
)

func Receiver(v interface{}) {

	switch v.(type) {
	case string:
		fmt.Println("This is a string")
	case float64:
		fmt.Println("This is a float 64")
	case float32:
		fmt.Println("This is a float 32")
	case bool:
		fmt.Println("This is a bool")
	case int:
		fmt.Println("This is a Int")
	}

}

type Player struct {
	name string

	age int
}

type Players []Player

func (p Players) Len() int {
	return len(p)
}

func (p Players) Less(i, j int) bool {
	return p[i].age > p[j].age
}

func (p Players) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	players := Players{{"测试1", 1}, {"测试2", 2}, {"测试3", 3}}

	Receiver("测试")
	Receiver(123)
	Receiver(123.00)
	Receiver(float32(123.00))
	Receiver(true)

	fmt.Println(players)

	sort.Sort(players)

	fmt.Println(players)

}
