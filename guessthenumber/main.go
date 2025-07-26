package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	r := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(r)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Guess the random number")
	target := randomNumber.Intn(10) + 1
	fmt.Println(target)
	for scanner.Scan() {
		if input, err := strconv.Atoi(scanner.Text()); err != nil {
			fmt.Println("error converting input please make sure its a number between 1-10")
			continue
		} else if input == target {
			fmt.Println("you win")
			break
		} else {
			fmt.Println("try again...:(")
			if input > target {
				fmt.Println("too high")
			} else if input < target {
				fmt.Println("too low")
			}
		}
	}

}
