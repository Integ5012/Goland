package main

import "fmt"

/* comment
   ==================================
   GROUP Ewan : BARANGAY AYUDA DISTRIBUTION SYSTEM
   ==================================
*/

// CONSTANTS: di nagbabago guys habang nag rurun yung program
const (
	MinAge            = 18 // minimum qualified sa age
	InitialSupplyEach = 3  // ayuda packs given per purok kunwari tatlo lang yung available
)

// STRUCT: mga kaylangan na info per resident para tatawagin nalang tong method
type Resident struct {
	Name       string // DATA TYPE: string
	Age        int    // DATA TYPE: int
	Purok      int
	HasClaimed bool // DATA TYPE: bool
}

// MAPS: Purok number -> Purok name
var purokNames = map[int]string{
	1: "Purok Nico",
	2: "Purok Jenny",
	3: "Purok NatNat",
}

// ARRAYS: fixed-size (3 puroks, hindi nagbabago yung laki)
// index 0 = Purok 1, index 1 = Purok 2, index 2 = Purok 3
var supply = [3]int{InitialSupplyEach, InitialSupplyEach, InitialSupplyEach}

//JAVA COMPARISON
//int[] supply = {3, 3, 3};

var originalSupply = supply

//JAVA COMPARISON
//int[] originalSupply = supply; // this copies the REFERENCE, not the data

// FUNCTIONS + SWITCH: map a purok number
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

// what age sila?
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

// FUNCTIONS + CONDITIONS + OPERATORS: check if a resident or hindi resident
// currently qualifies for ayuda.
func isEligible(r Resident) bool {
	hasSupply := supply[r.Purok-1] > 0 // comparison operator: >
	isAdult := r.Age >= MinAge         // comparison operator: >=
	notYetClaimed := !r.HasClaimed     // logical operator: !
	return isAdult && hasSupply && notYetClaimed
}

// FUNCTIONS: process one resident's claim.
// Takes a POINTER (*Resident) so the change actually sticks
// on the real entry inside the slice, not just a local copy.
func processClaim(r *Resident) bool {
	if !isEligible(*r) {
		return false
	}
	r.HasClaimed = true
	supply[r.Purok-1]-- // arithmetic operator: decrement
	return true
}

// FUNCTIONS:
// eto yung reason, kung bakit rejected sila
func rejectionReason(r Resident) string {
	switch {
	case r.Age < MinAge:
		return "underage"
	case r.HasClaimed:
		return "already claimed"
	case supply[r.Purok-1] <= 0:
		return "no more supply"
	default:
		return "unknown"
	}
}

func main() {
	// OUTPUT
	fmt.Println("=== Barangay Ayuda Distribution System ===")

	// VARIABLES: two styles - explicit "var" and short ":="
	var totalClaimed int = 0 // explicit declaration with type
	totalRejected := 0       // short declaration, type inferred

	// SLICES: dynamic list of residents pwedeng mag dagdag basta append nalang
	residents := []Resident{
		{Name: "Ramel", Age: 45, Purok: 1},
		{Name: "Lex", Age: 16, Purok: 1},                    // rejected: underage
		{Name: "Gian", Age: 40, Purok: 1, HasClaimed: true}, // rejected: already claimed
		{Name: "Zymon", Age: 38, Purok: 1},
		{Name: "Chollo", Age: 55, Purok: 1},
		{Name: "Von", Age: 33, Purok: 1}, //rejected: no more supply
		{Name: "Kisha", Age: 25, Purok: 2},
		{Name: "Kasima", Age: 68, Purok: 2},
		{Name: "Luis", Age: 30, Purok: 3},
		{Name: "Babes", Age: 50, Purok: 3},
		{Name: "Shyla", Age: 5, Purok: 3}, //rejected: underage
	}

	/*JAVA COMPARISON
	ArrayList<Resident> residents = new ArrayList<>(List.of(
		new Resident("Ramel", 45, 1, false),
		new Resident("Lex", 16, 1, false)
	));
	*/

	fmt.Println("\n--- Slice Info Before Walk-ins ---")
	fmt.Printf("residents: len=%d, cap=%d\n", len(residents), cap(residents))
	fmt.Println("\n--- Processing Claims ---")

	walkIns := make([]Resident, 0, 2)
	walkIns = append(walkIns, Resident{Name: "Arnold", Age: 38, Purok: 2})
	walkIns = append(walkIns, Resident{Name: "Tunying", Age: 86, Purok: 1})

	//JAVA COMPARISON
	// ArrayList<Resident> walkIns = new ArrayList<>(2); // sets initial capacity only

	residents = append(residents, walkIns...)

	fmt.Println("\n--- Slice Info After Walk-ins Appended ---")
	fmt.Printf("residents: len=%d, cap=%d (tumaas dahil sa append)\n", len(residents), cap(residents))

	purok1Residents := residents[0:6]
	fmt.Println("\n--- Purok 1 Residents Only (via slicing) ---")
	for _, r := range purok1Residents {
		fmt.Println(" -", r.Name)
	}

	purok1Residents[0].HasClaimed = true //JAVA COMPARISON -> purok1Residents.get(0).setHasClaimed(true);
	fmt.Println("\n--- Shared Memory Check ---")
	fmt.Printf("residents[0] (%s) HasClaimed = %v -- nagbago kahit sa purok1Residents lang ito ginawa\n",
		residents[0].Name, residents[0].HasClaimed)
	purok1Residents[0].HasClaimed = false

	fmt.Println("\n--- Processing Claims ---")

	// LOOPS: classic for loop, index-based
	for i := 0; i < len(residents); i++ {
		r := &residents[i]
		group := ageGroup(r.Age) // check kung minor,
		zone := purokZone(r.Purok)

		// CONDITIONS: if / else branching sa claim result
		if processClaim(r) {
			totalClaimed++ // arithmetic operator: increment
			fmt.Printf("[CLAIMED]  %-12s | %-14s | %s (%s) | %d packs left\n",
				r.Name, group, purokNames[r.Purok], zone, supply[r.Purok-1])
		} else {
			totalRejected++
			reason := rejectionReason(*r)
			fmt.Printf("[REJECTED] %-12s | %-14s | %s (%s) | reason: %s\n",
				r.Name, group, purokNames[r.Purok], zone, reason)
		}
	}

	// OUTPUT: summary report
	fmt.Println("\n--- Summary Report ---")
	fmt.Println("Total Residents:", len(residents))
	fmt.Println("Total Claimed:  ", totalClaimed)
	fmt.Println("Total Rejected: ", totalRejected)

	// DATA TYPES + OPERATORS: float64 and type conversion,
	// using the arithmetic operators + and /
	var totalAge float64 = 0
	for _, r := range residents { // LOOPS: for-range form parang for each
		totalAge += float64(r.Age) // type conversion int -> float64
	}
	averageAge := totalAge / float64(len(residents))
	fmt.Printf("Average Age of Residents: %.2f\n", averageAge)

	// ARRAYS + MAPS + LOOPS: final remaining supply per purok
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
