package main

const (
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
)

func main() {

	// go constants are immutable and evaluated at compile time

	// declaring constants

	const Pi = 3.14 // untyped constant no explicit type
	// untyped constants adapt there type based on usage

	const Greeting string = "Hi." // typed constant has a fixed type

	// const must be initalised with values that can be determined at compile time
	//  runtime evaluated values cannot be assigned to constants
	// constants have no short declaration

}
