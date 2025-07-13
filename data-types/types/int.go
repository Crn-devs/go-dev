package types

import (
	"fmt"
)

func DescribeIntType() {

	fmt.Println("This is the signed integer type")
	fmt.Println("We have 5 different keywords for integers in go")
	fmt.Println("int: size can range from -2147483648 to 2147483647 in 32 bit systems. -9223372036854775808 to 9223372036854775807 in 64 bit systems")
	fmt.Println("int8: 8bits 1byte size can range from -128 to 127")
	fmt.Println("int16: 16bits 2bytes size can vary from -32768 to 32767")
	fmt.Println("int32: 32bits 4bytes size can vary from -2147483648 to 2147483647")
	fmt.Println("int64: 64bits 8bytes size can vary from -9223372036854775808 to 9223372036854775807")
}
