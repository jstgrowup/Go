package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to files")
	content := "file content inside the file - Subham"
	file, err := os.Create("./subham.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is",length)
	file.Close()
}
