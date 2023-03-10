package main

import "fmt"

func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("threee")
	fmt.Println("hello")
	// hello
	// threee
	// two
	// one
	// LIFO principle
	myDefeer()
	// hello
	// 4
	// 3
	// 2
	// 1
	// 0
	// threee
	// two
	// one
}
func myDefeer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
