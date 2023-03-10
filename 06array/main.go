package main

import "fmt"

func main() {
	fmt.Println("welcome")
	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "Peach"
	fmt.Println(len(fruitlist))
	var vegList = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("vegy luist", vegList, len(vegList))
}
