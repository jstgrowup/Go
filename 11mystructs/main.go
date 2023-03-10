package main

import "fmt"

func main() {
	fmt.Println()
	// no inheritance in golang no super or parent
	subham := User{"subham", "deysubham999@gmail.com", true, 23}
	fmt.Println(subham)
	fmt.Printf("hitesh details are: %v\n",subham)
	fmt.Printf("hitesh details are: %v\n",subham.Name)
	fmt.Printf("hitesh details are: %v\n",subham.Age)


}

type User struct {
	Name   string
	Email  string
	States bool
	Age    int
}
