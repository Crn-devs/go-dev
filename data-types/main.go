package main

import (
	"bufio"
	"data-types/types"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Please Select A Data Type To Research")
	scanner := bufio.NewScanner(os.Stdin)

	menu()

	for scanner.Scan() {
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Please select from the list, error on your input")
			menu()
		}

		switch choice {
		case 1:
			types.DescribeIntType()
			menu()
		case 2:
			types.DescribeFloatType()
			menu()
		case 3:
			types.DescribeBoolType()
			menu()
		case 4:
			types.DescribeStringType()
			menu()
		case 5:
			types.DescribeArrayType()
			menu()
		case 6:
			os.Exit(0)

		default:
			fmt.Println("Not an option")
		}

	}

}

func menu() {

	fmt.Println(`
		1: integers
		2: floats
		3: bool
		4: string
		5: array
		6: Exit
	`)
}
