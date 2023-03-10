package main

import "fmt"

func main() {
	// map is kindof like an object or you can say a key value pair
	// map[string]string the string that is inside the bracket is the key and the second one is the value

	languages := make(map[string]string)
	languages["js"] = "Javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	fmt.Println("list of all", languages)

	fmt.Println(languages)
	delete(languages, "rb")
	fmt.Println("deleted or updated map", languages)
	// loops are interesting in go
	for key, value := range languages {
		fmt.Println("for key %v\n", key, value)
	}
	
}
