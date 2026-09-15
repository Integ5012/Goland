package main

// IMPORTS: extra packages we need besides fmt.
// bufio  -> Scanner, so we can read a WHOLE line (names with spaces)
// os     -> os.Stdin is the keyboard
// strconv -> convert typed text like "25" into an int
// strings -> TrimSpace / ToLower for cleaning input
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* comment
   ==================================
   GROUP Ewan : BARANGAY AYUDA DISTRIBUTION SYSTEM
   Interactive version -- the user types residents, then claims are processed.
   ==================================
*/

// CONSTANTS: di nagbabago guys habang nag rurun yung program
const (
	MinAge            = 18 // minimum qualified sa age
	InitialSupplyEach = 3  // ayuda packs given per purok kunwari tatlo lang yung available
)

// STRUCT: mga kaylangan na info per resident para tatawagin nalang tong type
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

// JAVA COMPARISON
// int[] supply = {3, 3, 3};

var originalSupply = supply

// JAVA COMPARISON
// int[] originalSupply = supply; // this copies the REFERENCE, not the data

// INPUT: one Scanner shared by the whole program.
// NewScanner(os.Stdin) = read from the keyboard, not from a file.
var input = bufio.NewScanner(os.Stdin)

// FUNCTIONS + SWITCH: map a purok number to its zone
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

// FUNCTIONS + tagless SWITCH: what age group sila?
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

// FUNCTIONS + CONDITIONS + OPERATORS: check if a resident qualifies for ayuda.
func isEligible(r Resident) bool {
	// guard: invalid purok would panic if we index supply[r.Purok-1]
	if r.Purok < 1 || r.Purok > len(supply) {
		return false
	}
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

// FUNCTIONS: eto yung reason, kung bakit rejected sila
func rejectionReason(r Resident) string {
	switch {
	case r.Purok < 1 || r.Purok > len(supply):
		return "invalid purok"
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

// readLine prints a prompt then waits for the user to press Enter.
// Scan() reads until newline, so "Ana Cruz" is one name, not two tokens.
func readLine(prompt string) string {
	fmt.Print(prompt)
	if !input.Scan() {
		return "" // Ctrl+Z / end of input
	}
	// TrimSpace removes leading/trailing spaces and the leftover newline
	return strings.TrimSpace(input.Text())
}

// readInt reuses readLine, then strconv.Atoi converts "25" -> 25.
// error is returned if the user typed letters instead of a number.
func readInt(prompt string) (int, error) {
	text := readLine(prompt)
	return strconv.Atoi(text)
}

func printMenu() {
	fmt.Println("\n========== MENU ==========")
	fmt.Println("1. Register a resident")
	fmt.Println("2. Register a walk-in")
	fmt.Println("3. View residents")
	fmt.Println("4. Process all claims")
	fmt.Println("5. View remaining supply")
	fmt.Println("6. View summary report")
	fmt.Println("7. Load sample residents")
	fmt.Println("8. Exit")
	fmt.Println("==========================")
}

// registerResident asks for Name, Age, Purok, and HasClaimed.
// POINTER to the slice (*[]Resident) so append updates the list in main,
// not a copy of the slice header.
func registerResident(residents *[]Resident, label string) {
	fmt.Println("\n--- Register", label, "---")
	name := readLine("Name: ")
	if name == "" {
		fmt.Println("Name cannot be empty.")
		return // early return: stop this function, go back to the menu
	}

	age, err := readInt("Age: ")
	if err != nil || age < 0 {
		fmt.Println("Invalid age.")
		return
	}

	fmt.Println("Purok options: 1 = Purok Nico, 2 = Purok Jenny, 3 = Purok NatNat")
	purok, err := readInt("Purok (1-3): ")
	if err != nil {
		fmt.Println("Invalid purok.")
		return
	}
	// comma-ok map lookup: ok is false if the key is missing
	if _, ok := purokNames[purok]; !ok {
		fmt.Println("Purok must be 1, 2, or 3.")
		return
	}

	alreadyClaimed := false
	claimedAnswer := strings.ToLower(readLine("Already claimed? (y/n): "))
	if claimedAnswer == "y" || claimedAnswer == "yes" {
		alreadyClaimed = true
	}

	// append grows the slice; *residents writes back to main's variable
	*residents = append(*residents, Resident{
		Name:       name,
		Age:        age,
		Purok:      purok,
		HasClaimed: alreadyClaimed,
	})
	fmt.Printf("Added %s (age %d, %s). Slice len=%d, cap=%d\n",
		name, age, purokNames[purok], len(*residents), cap(*residents))
}

func viewResidents(residents []Resident) {
	fmt.Println("\n--- Residents ---")
	if len(residents) == 0 {
		fmt.Println("No residents registered yet.")
		return
	}
	// LOOPS: for-range; i is index, r is a COPY of each Resident
	for i, r := range residents {
		fmt.Printf("%2d. %-12s | age %3d | %-12s | claimed=%v | %s\n",
			i+1, r.Name, r.Age, purokNames[r.Purok], r.HasClaimed, ageGroup(r.Age))
	}
	fmt.Printf("slice: len=%d, cap=%d\n", len(residents), cap(residents))
}

func processAllClaims(residents []Resident) (int, int) {
	fmt.Println("\n--- Processing Claims ---")
	if len(residents) == 0 {
		fmt.Println("No residents to process.")
		return 0, 0
	}

	// VARIABLES: two styles - explicit "var" and short ":="
	var totalClaimed int = 0 // explicit declaration with type
	totalRejected := 0       // short declaration, type inferred

	// LOOPS: classic for loop, index-based so we can take a pointer
	for i := 0; i < len(residents); i++ {
		r := &residents[i] // POINTER to the real slice element
		group := ageGroup(r.Age)
		zone := purokZone(r.Purok)

		// CONDITIONS: if / else branching sa claim result
		if processClaim(r) {
			totalClaimed++ // arithmetic operator: increment
			fmt.Printf("[CLAIMED]  %-12s | %-14s | %s (%s) | %d packs left\n",
				r.Name, group, purokNames[r.Purok], zone, supply[r.Purok-1])
		} else {
			totalRejected++
			reason := rejectionReason(*r) // *r dereference: pointer -> value
			fmt.Printf("[REJECTED] %-12s | %-14s | %s (%s) | reason: %s\n",
				r.Name, group, purokNames[r.Purok], zone, reason)
		}
	}
	return totalClaimed, totalRejected
}

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

func viewSummary(residents []Resident, lastClaimed, lastRejected int) {
	fmt.Println("\n--- Summary Report ---")
	fmt.Println("Total Residents:", len(residents))
	fmt.Println("Last batch Claimed: ", lastClaimed)
	fmt.Println("Last batch Rejected:", lastRejected)

	if len(residents) == 0 {
		return
	}

	// DATA TYPES + OPERATORS: float64 and type conversion
	var totalAge float64 = 0
	for _, r := range residents { // _ ignores the index
		totalAge += float64(r.Age) // type conversion int -> float64
	}
	averageAge := totalAge / float64(len(residents))
	fmt.Printf("Average Age of Residents: %.2f\n", averageAge)
}

func sampleResidents() []Resident {
	// SLICES: dynamic list -- pwedeng mag dagdag later with append
	return []Resident{
		{Name: "Ramel", Age: 45, Purok: 1},
		{Name: "Lex", Age: 16, Purok: 1},                    // rejected: underage
		{Name: "Gian", Age: 40, Purok: 1, HasClaimed: true}, // rejected: already claimed
		{Name: "Zymon", Age: 38, Purok: 1},
		{Name: "Chollo", Age: 55, Purok: 1},
		{Name: "Von", Age: 33, Purok: 1}, // rejected: no more supply
		{Name: "Kisha", Age: 25, Purok: 2},
		{Name: "Kasima", Age: 68, Purok: 2},
		{Name: "Luis", Age: 30, Purok: 3},
		{Name: "Babes", Age: 50, Purok: 3},
		{Name: "Shyla", Age: 5, Purok: 3}, // rejected: underage
	}
}

func main() {
	// OUTPUT
	fmt.Println("=== Barangay Ayuda Distribution System ===")
	fmt.Println("Type a menu number, then press Enter.")

	/* JAVA COMPARISON
	ArrayList<Resident> residents = new ArrayList<>();
	Scanner sc = new Scanner(System.in);
	*/

	// SLICES: start empty; grow with append when the user registers people
	residents := make([]Resident, 0)
	lastClaimed, lastRejected := 0, 0

	// infinite loop: menu keeps showing until the user picks 8 (or return)
	for {
		printMenu()
		choice, err := readInt("Choice: ")
		if err != nil {
			fmt.Println("Please enter a number from 1 to 7.")
			continue // skip the rest of this loop, show the menu again
		}

		// SWITCH: one case per menu number
		switch choice {
		case 1:
			registerResident(&residents, "resident") // & = address of the slice

		case 2:
			viewResidents(residents)
		case 3:
			lastClaimed, lastRejected = processAllClaims(residents)
		case 4:
			viewSupply()
		case 5:
			viewSummary(residents, lastClaimed, lastRejected)
		case 6:
			// ... unpacks the sample slice into individual elements for append
			residents = append(residents, sampleResidents()...)
			fmt.Printf("Sample residents loaded. len=%d, cap=%d\n", len(residents), cap(residents))
		case 7:
			fmt.Println("Salamat. Program ended.")
			return // leave main() = end the program
		default:
			fmt.Println("Please enter a number from 1 to 7.")
		}
	}
}
