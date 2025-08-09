package main

/**
	Slices in Go

	a slice is not an array its a descriptor that gives a view into an underlying array

	internally in go slices are a struct called a struct header

	they have the following properties
		pointer: pointer to the underlying array
		length: number of accessible elements
		capacity: max number of elements before hitting the arrays end

**/

/**
	Declaring and Initalising

	slice literal
		nums := []int{1,2,3}

	from array
		nums := [5]int{1,2,3,4,5}
		numSlice := nums[0:4]

	using make()
		nums := make([]string, 3) // initalised with 3 elements and a capacity of 3 all elements zero valued

		nums := make([]string, 3, 5) // initalised with 3 elements but capacity of 5. 3 elements zero valued
**/

/**
	Length vs Capacity

		the length refers to the number of elements inside the slice which are initalised
		the capacity refers to the number of elements the underlying array contains

		nums := []int{1,2} // length of nums is 2
						   // capacity is 2

		if we slice from an array then capacity refelects that

		nums := [5]int{1,2,3,4,5}
		s := nums[0:2]

		len(s) // 2 -> length is the amount of elements inside the slice [1,2]
		cap(s) // 5 -> capacity refers to the underlying arrays capacity

		if however we append to the s slice above 5 elements a new backing array will be created

		s = append(s, nums) // backing array will be updated as there are now 7 elements

**/

/**
	Appending and Capacity Behaviour

		nums := [5]int{1, 2, 3, 4, 5}
		s := nums[0:2]

		fmt.Println(len(s)) // 2 -> length is the amount of elements inside the slice [1,2]
		fmt.Println(cap(s)) // 5 -> capacity refers to the underlying arrays capacity

		//if however we append to the s slice above 5 elements a new backing array will be created

		for _, v := range nums {
			s = append(s, v)
			fmt.Println("cap on s is growing:", cap(s))
				// if we log the capacity on each iteration we can see it double from 5-10 once we append to a length 5>
		}

		fmt.Println(len(s)) // 7 -> length is the amount of elements inside the slice [1,2,1,2,3,4,5]
		fmt.Println(cap(s)) // 10 -> capacity typically doubles when the underlying array isnt large enough
							// it does this by creating a new backing array

**/

/**
	Copying slices

		copying can be done with the built in copy function

			src := []int{1, 2, 3}
			dst := make([]int, len(src))
			copy(dst, src)
			copy copies whichever slice is smaller out of len(dst) and len(src).
			if the dst is 4 but the src is 5 only 4 elements will be copied
			if the dst is 6 but the src is 3 only 3 will be copied over
**/

/**
	Removing slice elements

	no built in method or function

	best idiomatic solution is

	i := index to remove
	s = append(s[:i], s[i+1:]...)

	s is equal to the result of appending from the index
			 we want to remove and then appending from one after that index

**/

/**
	Cloning a slice
		copy to a new slice with a new backing array

		2 idiomatic options

		n := []int{1,2,3,4,5}

		clone := append([]int(nil), n...)

			the nil cast tells go create a nil slice, when we append to a nil slice
			go allocates a new array with exactly the needed capacity

		Equivalently

		clone := make([]int, len(s))
		copy(clone, s)

**/

/**
	Capacity leaks
		When a slice from a large array only references a small section
		the large array in memory cannot be cleaned up due to the slice holding a reference

		big := make([]byte, 1e6) // 1 million bytes
		small := big[:10]        // small slice of length 10

		// small keeps the big array alive in memory

	we can fix this by creating a new backing array and copying exactly what we need

	smallCopy := append([]int(nil), big[:10]...)

	smallCopy now only references those elements from index 10 onwards and has a capacity reflecting that

	full example:
			big := make([]int, 1e6/2)

			fmt.Println(len(big)) // 500,000
			fmt.Println(cap(big)) // 500,000

			kindaSmall := big[:len(big)/2]

			fmt.Println(len(kindaSmall)) // 250,000
			fmt.Println(cap(kindaSmall)) // still 500,000

			// if we clone though

			actuallySmall := append([]int(nil), big[:len(big)/2]...) // this line clones the large array values but not the backing array

			fmt.Println(len(actuallySmall)) // 250,000
			fmt.Println(cap(actuallySmall)) // now: 250,880 fixed the capacity leak

**/

/**
	Good to know

	Trim capacity leaks with copy or append(nil, slice...).

	Slices are “array pointers + length + capacity” — header copy not data copy.

	Passing a slice to a function copies the header, not the underlying data.

	Element changes inside a function are visible; slice length/cap changes are not unless you return or pass *[]T.

		a few built in slice methods from the slices package
			nums := []int{1, 2, 3}
			nums2 := []int{1, 2, 3, 4, 5}

			fmt.Println(slices.Equal(nums, nums2[:3])) // true

			fmt.Println(slices.Contains(nums2, 5)) // true

			fmt.Println(slices.Index(nums, 5)) // -1 not found
**/

func main() {

}
