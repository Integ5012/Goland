package main

import "fmt"

type Resident struct {
	Name       string
	Age        int
	Purok      int
	HasClaimed bool
}

var purokNames = map[int]string{
	1: "Purok Nico",
	2: "Purok Jenny",
	3: "Purok NatNat",
}

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

// SLICES: dynamic slice creation and manipulation
func sampleResidents() []Resident {
	return []Resident{
		{Name: "Ramel", Age: 45, Purok: 1},
		{Name: "Lex", Age: 16, Purok: 1},
		{Name: "Gian", Age: 40, Purok: 1, HasClaimed: true},
		{Name: "Zymon", Age: 38, Purok: 1},
		{Name: "Chollo", Age: 55, Purok: 1},
		{Name: "Von", Age: 33, Purok: 1},
		{Name: "Kisha", Age: 25, Purok: 2},
		{Name: "Kasima", Age: 68, Purok: 2},
		{Name: "Luis", Age: 30, Purok: 3},
		{Name: "Babes", Age: 50, Purok: 3},
		{Name: "Shyla", Age: 5, Purok: 3},
	}
}

func viewResidents(residents []Resident) {
	fmt.Println("\n--- Residents List (Slice) ---")
	if len(residents) == 0 {
		fmt.Println("No residents registered yet.")
		return
	}
	for i, r := range residents {
		fmt.Printf("%2d. %-12s | age %3d | %-12s | claimed=%v | %s\n",
			i+1, r.Name, r.Age, purokNames[r.Purok], r.HasClaimed, ageGroup(r.Age))
	}
	fmt.Printf("slice metadata: len=%d, cap=%d\n", len(residents), cap(residents))
}

func main() {
	// Demonstrating Slice Operations
	fmt.Println("=== SLICES DEMO ===")
	residents := sampleResidents()
	viewResidents(residents)

	// Appending to dynamic slice
	newResident := Resident{Name: "Nico", Age: 22, Purok: 1, HasClaimed: false}
	residents = append(residents, newResident)

	fmt.Println("\n[After Appending New Resident]")
	viewResidents(residents)
}
