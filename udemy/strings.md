all strings are mutable in golang
you can change the value of strings again and again
    // var m1 string
	m1 := "hello"
    // m1:="world" 
	// this will throw an error because initialization of a variable is only allowed once not twice 
	m1 = "hello world"
	// this will not throw an error because reassignment is allowed
	fmt.println(m1)



	// var m1 string
	// m1 := "hello"

    m1:="world" 
	// this will throw an error because initialization of a variable is only allowed once not twice 
	m1 = "hello world"
	// this will not throw an error because reassignment is allowed
	fmt.Println(m1)
    p1:="hello world"
	p2:="world"
    
	fmt.Println(strings.Contains(p1, p2))