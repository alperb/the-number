package validator

import "math"

func Validate(guess int64) bool {
	if checkIfNumberIsEven(guess) {
		if !checkIfNumberIsCoPrimeWith42(guess) {
			if checkIfNumberExistsInFibonacciSeries(guess) {
				if checkIfNumberHasExactly7Digits(guess) {
					return true
				}
			}
		}
	}

	return false
}

func checkIfNumberIsEven(guess int64) bool {
	if guess%2 == 0 {
		return true
	} else {
		return false
	}
}

func checkIfNumberIsCoPrimeWith42(guess int64) bool {
	if guess%42 == 0 {
		return true
	} else {
		return false
	}
}

func checkIfNumberExistsInFibonacciSeries(guess int64) bool {
	firstTerm := (5 * (guess * guess)) + 4
	secondTerm := (5 * (guess * guess)) - 4

	if checkIfNumberIsPerfectSquare(firstTerm) || checkIfNumberIsPerfectSquare(secondTerm) {
		return true
	} else {
		return false
	}
}

func checkIfNumberIsPerfectSquare(number int64) bool {
	squareRoot := int(math.Sqrt(float64(number)))
	if int64(squareRoot*squareRoot) == number {
		return true
	} else {
		return false
	}
}

func checkIfNumberHasExactly7Digits(guess int64) bool {
	if guess >= 1000000 && guess <= 9999999 {
		return true
	} else {
		return false
	}
}
