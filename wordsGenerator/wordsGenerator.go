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
	if noUpper {
		chars = lowercase + numbers
	} else if noLower {
		chars = uppercase + numbers
	} else if noNums {
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
	} else if special {
		chars += specialWords
	}
	fmt.Printf(" length: %v \n  special: %v \n", length, special)
	fmt.Printf(" noUpperCaseWords: %v \n", noUpper)
	fmt.Printf(" noLowerCaseWords: %v \n", noLower)
	fmt.Printf(" noNumbers: %v \n", noNums)
	fmt.Printf(" fullShow: %v \n", full)
	return chars
}
