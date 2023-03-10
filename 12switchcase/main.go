package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {
	fmt.Println("switch case")
	// rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("value of dice is ", dicenumber)
	switch dicenumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move to 3 spot")
	case 4:
		fmt.Println("you can move to 4 spot")
	case 5:
		fmt.Println("you can move to 5 spot")
	case 6:
		fmt.Println("you can move to 6 spot")
	default:
		fmt.Println("it is default")
	}
}
