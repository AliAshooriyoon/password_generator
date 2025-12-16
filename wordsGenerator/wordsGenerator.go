package wordsgenerator

import (
	"github.com/fatih/color"
)

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
	} else if noLower && !noUpper && !noNums {
		chars = uppercase + numbers
	} else if noNums && !noUpper && !noLower {
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
	color.Cyan(" length %v \n", length)
	color.Blue(" special: %v \n", special)
	color.Green(" noUpperCaseWords: %v \n", noUpper)
	color.Magenta(" noLowerCaseWords: %v \n", noLower)
	color.Yellow(" noNumbers: %v \n", noNums)
	color.Red(" fullShow: %v \n", full)
	return chars
}
