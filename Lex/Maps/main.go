package main

import "fmt"

/* ================================================================
   CLASS 3 of 4 — MAPS  (Barangay Ayuda)
   A map is a key -> value table (Java: HashMap).

   Run:  go run ./Lex/Maps
   ================================================================ */

func main() {
	fmt.Println("=== Barangay Ayuda — MAPS ===")

	// MAP: purok number -> purok display name
	purokNames := map[int]string{
		1: "Purok Nico",
		2: "Purok Jenny",
		3: "Purok NatNat",
	}

	fmt.Println("Lookup purok 2:", purokNames[2])

	// comma-ok: found is false when the key does not exist
	name, found := purokNames[9]
	fmt.Println("Lookup purok 9:", name, "| exists?", found)

	// MAP: resident name -> already claimed?
	claimed := map[string]bool{
		"Gian":  true,
		"Jenny": false,
		"Lex":   false,
	}

	fmt.Println("\nHas Gian claimed?", claimed["Gian"])
	fmt.Println("Has Jenny claimed?", claimed["Jenny"])

	// add / update an entry (walk-in)
	claimed["Arnold"] = false
	fmt.Println("After walk-in Arnold:", claimed)

	fmt.Println("\nAll puroks:")
	for num, label := range purokNames {
		fmt.Println(" ", num, "->", label)
	}
}
