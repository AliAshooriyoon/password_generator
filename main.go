package main

import (
	"fmt"

	"example.com/app/passGenerator"
)

func main() {
	pass := passgenerator.PassGenerator()
	fmt.Println("-----------------pass----------------------")
	fmt.Println(pass)
}
