package types

import "fmt"

func DescribeArrayType() {
	fmt.Println("This is the array type")
	fmt.Println("An array is a numbered sequence of elements of the same type")
	fmt.Println("The number of elements is called the length of the array")
	fmt.Println("Heres an example")
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T: has a length of %v and elements %v\n", nums, len(nums), nums)
}
