package main

import "fmt"

func main() {
	var num1, num2 int = 22, 3

	fmt.Println("inital numbers are", num1, num2)
	// basics
	fmt.Println("----Addition----")
	fmt.Println(num1 + num2)
	fmt.Println("----Addition----")

	fmt.Println("----Subtraction----")
	fmt.Println(num1 - num2)
	fmt.Println("----Subtraction----")

	fmt.Println("----Multiplication----")
	fmt.Println(num1 * num2)
	fmt.Println("----mulitiplication----")

	// problems arise in division
	// if both variables are integer type the result of the calculation is integer aswell
	// this means that any Decimal part (remainders) are truncted this leads to losing precision
	fmt.Println("----Division----")
	fmt.Println(num1 / num2)
	fmt.Println("----Division----")

	// even if we assign the result to a float the result is still truncated
	var result float64 = 22 / 7
	fmt.Println("result of 2 integers (22) / (7) divided, assigned to float64. result:", result)
	// as we can see the result is still truncated
	// To preserve precision we need atleast one operand to be of type float64
	// this forces the operation to be floating point division
	var num3 float64 = 22 / 7.0
	fmt.Println(num3) // correctly displayed

	// overflow example:

	var overflow uint8 = 255
	fmt.Println(overflow + 1) // result is 0 uint8 can only store 0-255
	var underflow uint8 = 0
	fmt.Println(underflow - 1) // result is 255 underflow

}

// Basic Arithmetic operators
/**

Addition: +
Subtraction: -
Division: /
Remainder(modulus): %

**/

// Operator precedence
/**

1: Parentheses ()
2: Multiplication, Division, Remainder
3: Addition, Subtraction

**/

// overflow and underflow
/**

overflow: when a value excedes the maximum limit of a type

underflow: when a value goes below the minimum limit for a type

example: uint8 can store from 0-255
		  if we assign a uint8 to 255 and add 1 it will overflow back to 0

**/

/**
Integers
Type	Size	Range
int8	8-bit	-128 to 127
uint8	8-bit	0 to 255
int16	16-bit	-32,768 to 32,767
uint16	16-bit	0 to 65,535
int32	32-bit	-2,147,483,648 to 2,147,483,647
uint32	32-bit	0 to 4,294,967,295
int64	64-bit	~±9.2 quintillion
uint64	64-bit	0 to ~18.4 quintillion
**/
