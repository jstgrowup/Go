package main

import "fmt"

func main() {
	fmt.Println("welcome to loops")

	days := []string{"Sunday", "Tuesday", "wednesday", "friday", "saturday"}
	fmt.Println(days)
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	for i := range days {
		fmt.Println("second for loop", days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	roughValue := 1
	for roughValue < 10 {
		if roughValue == 5 {
			break
		}
		fmt.Println("value is:", roughValue)
		roughValue++
	}
	 
}
