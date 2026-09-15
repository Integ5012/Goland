package main

import "fmt"

/* ================================================================
   CLASS 1 of 4 — STRUCT  (Barangay Ayuda)
   Go has no "class". A struct is the data part of a Java class.

   Run:  go run ./Lex/Struct
   ================================================================ */

// STRUCT: one resident's ayuda record (like Java instance fields)
type Resident struct {
	Name       string
	Age        int
	Purok      int
	HasClaimed bool
}

func main() {
	fmt.Println("=== Barangay Ayuda — STRUCT ===")

	// Create two Resident values (composite literals)
	ramel := Resident{
		Name:       "Ramel",
		Age:        45,
		Purok:      1,
		HasClaimed: false,
	}
	lex := Resident{Name: "Lex", Age: 16, Purok: 1, HasClaimed: false}

	fmt.Println("Whole struct:     ", ramel)
	fmt.Println("Field ramel.Name: ", ramel.Name)
	fmt.Println("Field ramel.Age:  ", ramel.Age)
	fmt.Println("Field ramel.Purok:", ramel.Purok)
	fmt.Println("HasClaimed:       ", ramel.HasClaimed)

	fmt.Println("\nSecond resident:", lex.Name, "| age", lex.Age)

	// Updating a field is like setting an instance variable
	ramel.HasClaimed = true
	fmt.Println("After ramel.HasClaimed = true:", ramel)
}
