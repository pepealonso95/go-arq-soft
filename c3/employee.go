package main

import (
	"fmt"
)

func Hey() {
	fmt.Printf("Hey")
}

type Employee struct {
	Id          int32
	FirstName   string
	LastName    string
	BadgeNumber int32
}
