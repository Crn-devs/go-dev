package main

import "fmt"

/**

	if statements allow us to change our code flow based on wether conditions result as true or false

	if condition {
		// code block
	}

	else statement is used when a if statements condition is unsatified

	if condition {
		// code block
	} else {
		// code block
	}

	if else can be chained over and over again to create long conditional checks where the last else clause
	is executed if the above conditions are not satisfied

		if condition {
		// code block
	} else if {
		// code block
	} else if {
		// code block
	} else


**/

// regular if else
func WeatherCheck(Temp float64) string {
	if Temp > 16.0 {
		return fmt.Sprintf("Today is a %v day", "warm")
	} else {
		return fmt.Sprintf("Today is a %v day", "cold")
	}
}

// mulit condition if else if
func ScoreCheck(score int) {
	if score == 100 {
		fmt.Println("Perfect Score A*")
	} else if score <= 90 {
		fmt.Println("Grade A")
	} else if score <= 80 {
		fmt.Println("Grade B")
	} else if score <= 50 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("fail")
	}
}

// nested if else statements
func IsGreaterThanButLessThan(number, limit, start int) {
	if number > start {
		if number < limit {
			fmt.Println("number is in bounds")
		} else {
			fmt.Println("number is out of bounds")
		}
	} else {
		fmt.Println("number is larger than the start range")
	}
}
