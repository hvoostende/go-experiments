package main

import (
	"fmt"

	"github.com/hvoostende/go-experiments/number-guess/random"
)

func main() {
	var name string
	var maxNumber int
	var guessedNumber int

	counter := 0

	// ask user to put in number
	fmt.Print("What's your name? ")
	fmt.Scan(&name)
	fmt.Printf("Let's play a game %v \n", name)
	fmt.Print("Let me think of a number below: ")
	fmt.Scan(&maxNumber)

	randomNumber := random.NewRandomNumber(maxNumber)

	// fmt.Println(randomNumber)
	// check if that's the correct one

	for randomNumber != guessedNumber {
		counter++
		fmt.Print("What number am I thinking of? ")
		fmt.Scan(&guessedNumber)

		if randomNumber == guessedNumber {
			fmt.Printf("Yes I was thinking of %v, you got it in %v tries!\n", randomNumber, counter)
		} else if guessedNumber < randomNumber {
			fmt.Printf("Sorry %v, it's higher\n", name)
		} else {
			fmt.Printf("Nope %v, it's lower\n", name)
		}

	}

}
