package main

import "fmt"

const InitialSupplyEach = 3

var purokNames = map[int]string{
	1: "Purok Nico",
	2: "Purok Jenny",
	3: "Purok NatNat",
}

func purokZone(purok int) string {
	switch purok {
	case 1:
		return "Zone A"
	case 2:
		return "Zone B"
	case 3:
		return "Zone C"
	default:
		return "Unknown Zone"
	}
}

// ARRAYS: fixed-size array state and operations
var supply = [3]int{InitialSupplyEach, InitialSupplyEach, InitialSupplyEach}
var originalSupply = supply

func viewSupply() {
	fmt.Println("\n--- Remaining Supply per Purok ---")
	for i := 0; i < len(supply); i++ {
		purokNum := i + 1
		fmt.Printf("%s (%s): %d packs left\n",
			purokNames[purokNum], purokZone(purokNum), supply[i])
	}
	fmt.Println("\n--- Array Copy Check (originalSupply vs supply) ---")
	fmt.Println("originalSupply (untouched array copy):", originalSupply)
	fmt.Println("supply (actual, modified array):       ", supply)
}

func main() {
	// Demonstrating Array Operations
	fmt.Println("=== ARRAYS DEMO ===")
	viewSupply()

	// Modify actual array to show copy behavior difference
	supply[0] = 0
	fmt.Println("\n[After modifying supply[0] = 0]")
	viewSupply()
}
