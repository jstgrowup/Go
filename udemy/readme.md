1. by default all the variables in Golang are initialized to 0
2. ```
package main
import "fmt"
func main(){
	//atomic variables
	var m1 int
	m1=2
	fmt.Println(m1)
	fmt.Println("hello world")
	// to declare multiple variables
	var (
		p1=2 
		p2=1
	)
	fmt.Println(p1+p2)
    // variable initialization
	m1:=3
	m2:=4
	
}
```