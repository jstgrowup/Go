package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome")
	var fruitList = []string{"apple", "tomato", "Peach"}
	fmt.Printf("type of printlist %T |\n", fruitList)
	fruitList = append(fruitList, "Mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	//  it will start from 1st index to the 2nd index
	fruitList = append(fruitList[:3])
	// this means that i dont have a default value just start from the 0th index and stop at 2nd index
	fmt.Println(fruitList)
	//  slices are deep down arrays only they are just syntactical sugars

	highScore := make([]int, 4)
	highScore[0] = 1
	highScore[1] = 34
	highScore[2] = 432
	highScore[3] = 343
	highScore[4] = 232
	highScore = append(highScore, 555, 666, 321)
	fmt.Println(highScore)
	sort.Ints(highScore)
	fmt.Println(highScore)
	// how to remove a element from the slice on the basis of index
	var courses = []string{"react", "swift", "python", "ruby", "JS"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	// op [react swift ruby JS]
	fmt.Println("courses",courses)

}
