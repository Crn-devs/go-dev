package main

import "fmt"

/**

Arrays in go are fixed length collections of the same element type

Arrays in go are fixed length: the size of the array is part of its type

value types: arrays are not pointers by default assigning or passing them copies the entire array

Homogeneous: all elements in an array must be of the same type

**/

/**

Declaration:

	Zero value declaration: var array = [3]string -> ["", "", ""]

	literal initalisation: array := [3]int{1, 2, 3} -> [1,2,3]

	partial initalisation: letters := [5]string{"A", "B"} -> ["A", "B", "", "", ""]

	inferred length: nums := [...]{1,2,3} -> length inferred as 3

	indexed initalisation: nums := [...]int{0: 1, 2: 2 } -> [1, 0, 2,]

**/

/**
	Length and indexing

	Built in length function
	len(array)

	array[0] indexing first element
	array[len(array) -1 ] indexing last element
**/

/**
	Arrays are value types

		arr := [3]int{1,2,3}

		arr2 := arr[1:2] // copys the array


		// passing an array to a function also copies it
		func change(arr [3]int) {
    		arr[0] = 42
		}

		nums := [3]int{1, 2, 3}
		change(nums)
		fmt.Println(nums) // [1 2 3]

	Modifying array values in place
		func changePtr(arr *[3]int) {
			arr[0] = 42
		}
		changePtr(&nums)
		fmt.Println(nums) // [42 2 3]

	**/

/**
	iterating through arrays

	// for range
	for i, v := range array {
		// code block
	}

	// treditional for loop

	for i := 0; i < len(nums); i++ {
    	fmt.Println(nums[i])
	}


**/

func main() {

	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	copiednums := nums

	copiednums[0] = 0 // no effect on previous array, array values are copied

	fmt.Println(copiednums) // 0,2,3,4,5,6,7,8,9,10
	fmt.Println(nums)       // 1,2,3,4,5,6,7,8,9,10

	var underlyingNums *[10]int

	underlyingNums = &nums
	underlyingNums[0] = 0 // changes through the pointer effect the original

	fmt.Println(underlyingNums)
	fmt.Println(nums) // changed by underlying nums as its a pointer to the array in memory not a copy

}
