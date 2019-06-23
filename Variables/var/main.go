package main

import "fmt"

// //Variable declartion example 1

// var (
// 	name string
// 	age int
// 	location string
// )

// //Variable declartion example 2
// var (
// name, location string
// age int
// )

// //Variable declartion example 3
// var name string
// var age int
// var location string

// //Declarations with initializers
// var (
// 	name string = "Prince Obeyrn"
// 	age int = 32
// 	location string = "Dorne"
// )


// //Inferred typing
// var (
// 	name = "Kyle"
// 	age = 32
// 	location = "Arizona"
// )

// //one line intializing
// var (
// 	name,location, age = "Ezra", "Los Angeles", 21
// )

func main()  {
	// assigning variables without var ()
	//Not available outside of keywords. Ex: var, func
	name,location := "Prince Oberyn", "Dorne"
	age := 32
	fmt.Printf("%s (%d) of %s", name, age, location)

	action := func(){
		//doing something
	}
	action()
}