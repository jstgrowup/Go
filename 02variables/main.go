package main

import "fmt"
const LoginToken string="sdafsd"
// in golang if the first keyword is capital than it will be considered as a public keyword

func main() {
	// fmt.Println("subha,")
	var username string = "subham"
	fmt.Println(username)
	fmt.Printf("variable is:%T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variabe %T \n", isLoggedIn)

	var anotherVariable int
	// if you declare an integer it will initiaize its value wiith 0
	fmt.Println(anotherVariable)
	fmt.Printf("dsf %T \n", anotherVariable)
	// implicit type
	var website = "subham"
	fmt.Println(website)
	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
