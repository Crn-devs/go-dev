package main

func main() {

	// naming conventions and cases

	// PascalCase (Captialised)
	// exported identifiers
	// "eg:" type User struct{}, func DoSomething()
	// structs, interfaces, enums

	// camelCase
	// unexported variables, functions, struct fields
	// "eg:", userId, func processData()

	// filenames
	// all lowercase
	// no hyphens because they are invalid in import paths
	// use an underscore for better readability
	// "eg:", utils.go, home_handler

	// Constants
	// CamelCase or PascalCase (not UPPERCASE)
	//   e.g., const MaxItems = 100

	// additional notes
	// Be consistent and follow Go idioms
	// keep variable names short and descriptive

	// package names should be all lowercase and have no case
	// godev, webapp

}
