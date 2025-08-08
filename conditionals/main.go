package main

import "fmt"

func main() {

	temperature := 24.0

	tempAsString := WeatherCheck(temperature)

	fmt.Println(tempAsString)

	scores := []int{43, 57, 86, 100, 81, 99, 94}

	for i := 0; i < len(scores); i++ {
		ScoreCheck(scores[i])
	}

	IsGreaterThanButLessThan(60, 100, 61)
	IsGreaterThanButLessThan(60, 100, 0)
	IsGreaterThanButLessThan(101, 100, 61)

	WeekdayOrWeekend(GetDay) // passing a func

	CheckType("hello")
	CheckType(50)
	CheckType(5.00)
	CheckType(-3.4e+38)

}
