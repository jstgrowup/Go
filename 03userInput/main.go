package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome"
	fmt.Println(welcome)
	// https://pkg.go.dev/ to download packages
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reader)

	// comma ok || err err syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for reading", input)
}
