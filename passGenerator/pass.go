package passgenerator

import (
	"crypto/rand"
	"flag"
	"math/big"

	wordsgenerator "example.com/app/wordsGenerator"
)

func PassGenerator() string {
	var length int
	var includesSpecial bool
	var includesNotUpper bool
	var includesNotLower bool
	var includesNotNumbers bool
	var full bool
	flag.IntVar(&length, "length", 0, " password length")
	flag.BoolVar(&includesSpecial, "special", false, "includes special")
	flag.BoolVar(&includesNotUpper, "no-upper", false, "includes not uppercase words")
	flag.BoolVar(&includesNotLower, "no-lower", false, "includes not lowercase words")
	flag.BoolVar(&includesNotNumbers, "nonumbers", false, "includes not numbers words")
	flag.BoolVar(&full, "full", true, "includes all characters")
	flag.Parse()
	password := make([]byte, length)

	chars := wordsgenerator.WordsGenerator(length, includesSpecial, includesNotUpper, includesNotLower, includesNotNumbers, full)

	lengthInt := big.NewInt(int64(len(chars)))
	for i := 0; i < length; i++ {
		lengthOfWords, _ := rand.Int(rand.Reader, lengthInt)
		password[i] = chars[lengthOfWords.Int64()]
	}
	return string(password)
}
