package main

import (
	"fmt"

	"example.com/app/passGenerator"
)

func main() {
	pass := passgenerator.PassGenerator()
	fmt.Println("-----------------password----------------------")
	fmt.Println(pass)
}
