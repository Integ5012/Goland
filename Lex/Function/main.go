package main

import "fmt"

/* ================================================================
   CLASS 2 of 4 — FUNCTIONS  (Barangay Ayuda)
   Behavior lives in functions (Java: methods / static helpers).

   Run:  go run ./Lex/Function
   ================================================================ */

type Resident struct {
	Name       string
	Age        int
	Purok      int
	HasClaimed bool
}

const MinAge = 18

func main() {
	fmt.Println("=== Barangay Ayuda — FUNCTIONS ===")

	gian := Resident{Name: "Gian", Age: 45, Purok: 1}
	lex := Resident{Name: "Lex", Age: 16, Purok: 1}
	nico := Resident{Name: "Nico", Age: 68, Purok: 2}

	// Call functions and print what they return
	show(gian)
	show(lex)
	show(nico)
}

func show(r Resident) {
	fmt.Printf("%-8s | %s | %s | eligible=%v | %s\n",
		r.Name,
		ageGroup(r.Age),
		purokZone(r.Purok),
		isEligible(r),
		rejectionReason(r),
	)
}

// FUNCTION: classify age for ayuda
func ageGroup(age int) string {
	switch {
	case age < 18:
		return "Minor"
	case age >= 60:
		return "Senior Citizen"
	default:
		return "Adult"
	}
}

// FUNCTION + SWITCH: purok number -> zone name
func purokZone(purok int) string {
	switch purok {
	case 1:
		return "Zone A (Purok Nico)"
	case 2:
		return "Zone B (Purok Jenny)"
	case 3:
		return "Zone C (Purok NatNat)"
	default:
		return "Unknown Zone"
	}
}

// FUNCTION: true only if adult and not yet claimed
func isEligible(r Resident) bool {
	return r.Age >= MinAge && !r.HasClaimed
}

// FUNCTION: human-readable why they cannot claim
func rejectionReason(r Resident) string {
	if r.Age < MinAge {
		return "reason: underage"
	}
	if r.HasClaimed {
		return "reason: already claimed"
	}
	return "reason: none (ok to claim)"
}
