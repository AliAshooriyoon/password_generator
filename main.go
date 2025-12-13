package main

import "fmt"

type (
	userIDType string
	userStruct struct {
		name     string
		age      int
		id       int
		national string
	}
)

func main() {
	users := []string{"ahmed", "ali"}
	fmt.Println(users[1])
	nameReturner := func(name string) string {
		fmt.Println("slm " + name)
		return "slm " + name
	}
	nameReturner("Ali")
}
