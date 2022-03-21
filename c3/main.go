package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Main Function")

	var employee = Employee{
		Id:          1,
		FirstName:   "First name",
		LastName:    "Last Name",
		BadgeNumber: 1000,
	}
	fmt.Printf(employee.FirstName)
}
