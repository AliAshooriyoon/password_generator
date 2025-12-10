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
	users := make(map[userIDType]userStruct)
	users["sara"] = userStruct{"Sara", 33, 1, "Israel"}
	users["ali"] = userStruct{"Ali", 18, 2, "Iran"}
	fmt.Println(users["ali"])
}
