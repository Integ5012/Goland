package main

import "fmt"

/* ================================================================
   CLASS 4 of 4 — POINTERS  (Barangay Ayuda)
   A pointer is the ADDRESS of a Resident.
   Use it so claim updates stick on the real record, not a copy.

   Run:  go run ./Lex/Pointers
   ================================================================ */

type Resident struct {
	Name       string
	Age        int
	Purok      int
	HasClaimed bool
}

func main() {
	fmt.Println("=== Barangay Ayuda — POINTERS ===")

	ramel := Resident{Name: "Ramel", Age: 45, Purok: 1, HasClaimed: false}

	fmt.Println("Start:                 ", ramel)

	// VALUE: Go copies the struct. Changes stay inside the function.
	claimByValue(ramel)
	fmt.Println("After claimByValue:    ", ramel, "  <- still false")

	// POINTER: &ramel = address. The function edits the original.
	claimByPointer(&ramel)
	fmt.Println("After claimByPointer:  ", ramel, "  <- now true")

	fmt.Println("\n&ramel is the address:", &ramel)
}

// Receives a COPY. Setting HasClaimed does not affect main's ramel.
func claimByValue(r Resident) {
	r.HasClaimed = true
	fmt.Println("  inside claimByValue: ", r)
}

// Receives *Resident (pointer). r.HasClaimed changes main's ramel.
func claimByPointer(r *Resident) {
	r.HasClaimed = true
	fmt.Println("  inside claimByPointer:", *r) // *r = follow the pointer
}
