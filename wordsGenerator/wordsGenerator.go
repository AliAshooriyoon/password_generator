package wordsgenerator

import "fmt"

const (
	lowercase    = "abcdefghijklmnopqrstuvwxyz"
	uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers      = "0123456789"
	specialWords = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func WordsGenerator(length int, special, noUpper, noLower, noNums, full bool) string {
	var chars string
	if noLower || noUpper || noNums {
		full = false
	}
	if noUpper {
		chars = lowercase + numbers
		fmt.Println("hier!")
	} else if noLower && !noUpper && !noNums {

		fmt.Println("hier2!")
		chars = uppercase + numbers
	} else if noNums && !noUpper && !noLower {

		fmt.Println("hier3!")
		chars = uppercase + lowercase
	}
	if noLower && noUpper {
		chars = numbers
	}
	if noNums && noLower {
		chars = uppercase
	}

	if noNums && noUpper {
		chars = lowercase
	}
	if special && full {
		chars += numbers + uppercase + lowercase + specialWords
	} else if full {
		return " "
	} else if special {
		chars += specialWords
	}
	fmt.Printf(" length: %v \n special: %v \n noUpperCaseWords: %v \n noLowerCaseWords: %v \n noNumbers: %v \n fullShow: %v \n", length, special, noUpper, noLower, noNums, full)
	return chars
}
