package main

import "fmt"

// Go
func main() {
	purok := 3

	switch purok {
	case 1, 2, 3, 4, 5, 6:
		fmt.Println("Zone A - High level")
	case 7, 8:
		fmt.Println("Zone B - Mid Level")
	case 9, 10:
		fmt.Println("Zone C - Lower Level")
	default:
		fmt.Println("Unknown Zone")
	}
}
