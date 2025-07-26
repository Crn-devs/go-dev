package main

import (
	"fmt"
	"strings"
)

func main() {

	// basic for loop

	for i := 0; i < 10; i++ {
		fmt.Println(i + 1) // prints one through ten
	}

	// interating over a collection with for loop and len()
	nums := []int{1, 5, 2, 54, 6, 7, 2}
	for i := 0; i < len(nums); i++ {
		println("index:", i, "Value", nums[i])
	}

	// iterating over a collection using built in range
	for i, v := range nums {
		fmt.Println("index", i, "value", v) // same result as before
	}

	// break for breaking out of loops

	for i := 0; i < len(nums); i++ {
		if nums[i] > 53 {
			break // break exits the loop it doesnt only exit the current iteration
		}
		fmt.Println("value", nums[i], "is not greater than 53")
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 53 {
			continue // continue skips the iteration of the loop to the next interation
		}
		fmt.Println("value", nums[i], "is not greater than 53")
	}

	printRightAngledTriangle(5)
	fmt.Println()
	printUpsideDownTriangle(5)
	fmt.Println()
	printCenteredTriangle(3)
	fmt.Println()
	printCenteredTriangleUpsideDown(3)
	fmt.Println()
	printDimondPattern(5)
}

// for loops
// for loops follow the for initalisation; condition; post; { CodeBlock }
// for i:=0; i<100; i++; {}
// initalisation is used to initalise the index value
// condition is evaluated before each loop iteration if its false the loop terminates
// post is evaluated after each interation; usually increments or decrements the counter value

// loops inclide a break and continue keyword
// break breaks out or stops the current loop execution
// continue skips the current loop execution block and moves onto the next

// loop pattern printing pratice
func printRightAngledTriangle(n int) {
	for i := range n {
		fmt.Println(strings.Repeat("*", i+1))
	}
}

// loop
// 5-0 = 5
// 5-1 = 4
// 5-2 = 3
// 5-3 = 2
// 5-4 = 1
// *****
// ****
// ***
// **
// *
func printUpsideDownTriangle(n int) {
	for i := range n {
		fmt.Println(strings.Repeat("*", n-i))
	}
}

// loop
// spaces: 5-0+1 6 total spaces Stars: 2*0+1 1 total star
// spaces: 5-1+1 5 total spaces Stars: 2*1+1 3 total star
// spaces: 5-2+1 4 total spaces Stars: 2*2+1 5 total star
// spaces: 5-3+1 3 total spaces Stars: 2*3+1 7 total star
// spaces: 5-4+1 1 total spaces Stars: 2*4+1 9 total star
//
//	    *
//	   ***
//	  *****
//	 *******
//	*********
func printCenteredTriangle(n int) {
	for i := range n {
		fmt.Println(strings.Repeat(" ", n-i+1) + strings.Repeat("*", 2*i+1))
	}
}

func printCenteredTriangleUpsideDown(n int) {
	for i := range n {
		fmt.Println(strings.Repeat(" ", i+1) + strings.Repeat("*", 2*(n-i)-1))
	}
}

//
//
//	*********
//	 *******
//	  *****
//	   ***
//	    *

// print dimond using only for loops no range

/**
	n = 5
output:
    *
   ***
  *****
   ***
    *
**/

/**
0 print 4 spaces 1 star n-1 for spaces or n-i+1, i+1 for stars
1 print 3 spaces 3 star n-2 for spaces or n-i+1, i+2 for stars
2 print 2 spaces 5 star n-3 for spaces or n-i+1, i+3 for stars
4 print 3 spaces 3 star n-2 for spaces or n-i+1, i+2 for stars
5 print 4 spaces 1 star n-1 for spaces or n-1+1, i+1 for stars
**/

func printDimondPattern(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(strings.Repeat(" ", n-i-1) + strings.Repeat("*", 2*i+1))
	}
	for i := 0; i < n; i++ {
		fmt.Println(strings.Repeat(" ", i) + strings.Repeat("*", 2*(n-i)-1))
	}
}
