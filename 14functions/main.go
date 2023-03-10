package main

import "fmt"

func main() {
	fmt.Println("welcome to functions")
	greeter()
	// func greeterTwo(){
	// 	fmt.Println("another method")
	// }
	// you are not allowed to write functions insode function
	//    how this function invokes without invocations
	result := adder(3, 4)
	fmt.Println("result", result)
	proResult := proAdder(2, 3, 4, 5)
	fmt.Println("proresult", proResult)
}
func adder(valOne int, valueTwo int) int {
	return valOne + valueTwo
}
func proAdder(values ...int) int {
	fmt.Println("hey i am values",values)
	// this will take all the values in the form of a slice and it 
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
func greeter() {
	fmt.Println("hello from golang")
}
