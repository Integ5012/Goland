package main

import "fmt"

// Go
func main() {
	purok := 1

	switch purok {
	case 1:
		fmt.Println("Zone A")
		fallthrough
	case 2:
		fmt.Println("Zone B")
	case 3:
		fmt.Println("Zone C")
	default:
		fmt.Println("Unknown Zone")
	}
}
