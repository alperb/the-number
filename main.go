package main

import (
	"fmt"
	"suctf/the-number/validator"
)

var welcomeMessage = `Howdy!
Would you like to play a game?
I'm thinking of a number between 1 and 1 trillion (1.000.000.000.000)
It's not a random number, but there are specific rules for how my number is made.
Can you guess it, even if I don't tell you the rules?`

func main() {
	fmt.Println(welcomeMessage)

	var guess int64

	fmt.Printf("Guess a number: ")
	fmt.Scanln(&guess)
	if validator.Validate(guess) {
		fmt.Println("You win!")
	} else {
		fmt.Println("Nope not it!")
	}
}
