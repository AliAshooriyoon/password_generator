package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Willkommen")
	var passLength int
	flag.IntVar(&passLength, "length", 0, "lenght of password")
	flag.Parse()
	fmt.Println(passLength)
}
