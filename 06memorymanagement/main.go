package main

//
import "fmt"

func main() {
	// pointers
	fmt.Println("welcome to pointer")

	var ptr *int
	// fmt.Println("value of pointer is", ptr)
	// if you dont provide a value than it will be nil
	myNumber := 23
	var realptr = &myNumber
	// & is also used to create a pointer but here the & means i am creating a pointer as well as referencing to the memory
	fmt.Println(realptr)
	fmt.Println("this is the pointer", ptr)
	fmt.Println("this is the memory allocation", *ptr)
	*ptr = *ptr * 2
	fmt.Println("new vlaue", myNumber)

}
