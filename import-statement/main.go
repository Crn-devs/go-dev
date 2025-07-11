package main

// unamed import
// named import
// go uses tree shaking when compiled to ensure only used functions are packaged in the
// binary and so any dead code isnt shipped
import (
	"fmt"
	foo "sort"
)

func main() {
	fmt.Println("The import statement")

	numbers := []int{3, 4, 6, 78, 9}

	// named import used
	foo.Slice(numbers, func(i, j int) bool {
		return i > j
	})

}
