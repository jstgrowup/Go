package main

import "fmt"

func main() {
	fmt.Println()
	// no inheritance in golang no super or parent
	subham := User{"subham", "deysubham999@gmail.com", true, 23}
	fmt.Println(subham)
	fmt.Printf("hitesh details are: %v\n", subham)
	fmt.Printf("hitesh details are: %v\n", subham.Name)
	fmt.Printf("hitesh details are: %v\n", subham.Age)

	User.GetStatus(User{})
	User.NewMail(User{})
	fmt.Printf("hitesh details are: %v\n", subham.Email)
}

type User struct {
	Name   string
	Email  string
	States bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user Active:", u.Email)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("new email is ", u.Email)
}
