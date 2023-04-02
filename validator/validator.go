package validator

import "math"

func Validate(guess int64) bool {
	if checkIfNumberIsOdd(guess) {
		if checkIfNumberIsCoPrimeWith37(guess) {
			if checkIfNumberExistsInFibonacciSeries(guess) {
				if checkIfNumberHasExactly8Digits(guess) {
					return true
				}
			}
		}
	}

	return false
}

func checkIfNumberIsOdd(guess int64) bool {
	if guess%2 != 0 {
		return true
	} else {
		return false
	}
}

func checkIfNumberIsCoPrimeWith37(guess int64) bool {
	if guess%37 == 0 {
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

func checkIfNumberHasExactly8Digits(guess int64) bool {
	if guess >= 10000000 && guess <= 99999999 {
		return true
	} else {
		return false
	}
}
