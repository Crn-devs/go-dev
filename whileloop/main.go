package main

import (
	"fmt"
)

// go doesnt have a native while loop

// we can use the for keyword to create a while loop
// we omit the initalisation and post statements to focus solely on the condition

func main() {

	// we can constrain the while loop based on a condition
	i := 0
	for i <= 20 {
		fmt.Println("i is:", i)
		i++
	}
	sum := 0
	// we can also create an infinite loop with the for keyword
	for {
		sum += 5
		if sum >= 50 {
			break
		}
		fmt.Println("sum is at", sum)
	}
	fmt.Println("sum reached 50")
	fmt.Println("ending infinite loop using break")

}
