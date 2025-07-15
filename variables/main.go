package main

import "fmt"

// default variable types (zero value)

// int 0
// bool false
// string ""
// pointers, slices, maps, functions are all assigned nil for zero value
// arrays and structs get all fields or elements initalised with zero values
//

func main() {

	// variable declarations

	// zero declaration
	var age int
	fmt.Println(age) // 0 integers zero value is 0

	// declare and instantiate
	var name string = "foo"
	fmt.Println(name) // foo

	// type infered from assigned value similar to below
	var lname = "bar"
	fmt.Println(lname)
	fmt.Printf("type of lname is %T\n", lname)

	// 	Short variable declarations - type is inferred
	message := "Hello, World."
	fmt.Println(message)
	fmt.Printf("type of message is %T\n", message)

	// nil variables
	var bar *int
	var numbers []int
	var users map[int]string
	var funky func()

	fmt.Println(bar == nil)     // true
	fmt.Println(numbers == nil) // true
	fmt.Println(users == nil)   // true
	fmt.Println(funky == nil)   // true

	// arrays and struct zero values
	type person struct {
		fName string
		lName string
	}

	var p1, pArr = person{}, [5]person{}

	fmt.Println(p1)   // zero's to person{fname: "", lName: ""}
	fmt.Println(pArr) // array of zero value persons

}
