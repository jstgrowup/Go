package main

import "fmt"

func main() {
	fmt.Println("if else")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watchout"
	} else {
		result = "nothing"
	}

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if num := 2; num < 10 {
		fmt.Println("num is less than 2")
	} else {
		fmt.Println("numn is greater than 2")
	}
	fmt.Println(result)
}
