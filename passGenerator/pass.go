package passgenerator

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	special   = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func PassGenerator() string {
	var length int
	flag.IntVar(&length, "length", 0, " password length")
	flag.Parse()
	password := make([]byte, length)
	chars := lowercase + uppercase + numbers
	lengthInt := big.NewInt(int64(len(chars)))
	for i := 0; i < length; i++ {
		lengthOfWords, _ := rand.Int(rand.Reader, lengthInt)
		password[i] = chars[lengthOfWords.Int64()]
		fmt.Println(i)
	}
	return string(password)
}
