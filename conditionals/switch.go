package main

import (
	"fmt"
	"time"
)

/**

	Switch cases offer one main benefit in programming which is readability

	they allow long if else statements to be condensed into a more readable manageable set of conditionals

	// switch condition (switch case default (fallthrough) )
	switch condition {
		case condition:
			// code block
		default:
			// code block executed if no conditions match the value
	}

	switch cases are much easier to scan compared to
	if else statements especially if we have more than 3+ cases


	switch cases in go are quite flexible and improved from c

	fallthrough bug: go has fixed the fallthrough bug and replaced that functionality with
	the fallthrough keyword which forces execution to continue into the next case regardless
	of wether the condition is satisfied.

	Value switch → switch someValue { case value1: ... }
	Expression (or condition) switch → switch { case condition1: ... }

	// type assertion

**/

func WeekdayOrWeekend(GetDay func() string) {

	// this example shows a switch case where we are evaluating the case based on the value after
	// the switch keyword and checking against that value also known as a value switch

	// using a function return value as a condition
	switch GetDay() { // the function must return the same type being checked
	case "Sunday", "Saturday":
		fmt.Println("Weekend")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	default:
		fmt.Println("not a weekday")
	}

}

func ApplesOrOranges(s string) {

	// this switch case checks the expression inside the case for each case
	// also known as an expression switch

	switch {
	case s == "apple":
		fmt.Println("apples")
	case s == "orange":
		fmt.Println("oranges")
	}
}

func GetDay() string {
	day := time.Now().Weekday()
	return day.String()
}

func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("string")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown type")
	}

}
