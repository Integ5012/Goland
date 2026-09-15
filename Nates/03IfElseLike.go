package main

import "fmt"

// Go
func main() {

	age := 10

	switch {
	case age < 18:
		fmt.Println("Minor")
	case age >= 60:
		fmt.Println("Senior Citizen")
	default:
		fmt.Println("Adult")
	}
}
