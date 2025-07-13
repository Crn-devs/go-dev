package types

import "fmt"

func DescribeBoolType() {

	fmt.Println("This is the Bool type")
	fmt.Println("bool types are used to store true or false")
	var truthy bool = true
	fmt.Printf("Type of variable truthy is %T and the value is %v", truthy, truthy)
}
