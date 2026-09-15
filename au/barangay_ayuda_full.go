package main

import "fmt"

/* ================================================================
   IT/CS 311 - PROGRAMMING PARADIGMS AND LANGUAGES
   GROUP EWAN -- GOLANG DEMO PROGRAM
   BARANGAY AYUDA DISTRIBUTION SYSTEM (Baguio City)
   ================================================================ */

// ---- Go Struct -------------------------------------------------
// STRUCT: groups related fields together, similar to a Java class
// but WITHOUT methods baked in -- methods are attached separately.
type Resident struct {
	Name       string // DATA TYPE: string
	Age        int    // DATA TYPE: int
	Purok      int    // DATA TYPE: int
	HasClaimed bool   // DATA TYPE: bool
}

// JAVA COMPARISON:
// class Resident {
//     String name; int age; int purok; boolean hasClaimed;
// }

// ---- Go Constants (also used by LUIS's section) -----------------
// CONSTANTS: fixed values, cannot change while the program runs.
const (
	MinAge            = 18 // minimum qualified age
	InitialSupplyEach = 3  // ayuda packs given per purok
)

// ---- Go Maps ------------------------------------------------------
// MAPS: key -> value lookup table. Here: purok number -> purok name.
var purokNames = map[int]string{
	1: "Purok Nico",
	2: "Purok Jenny",
	3: "Purok NatNat",
}

// ---- Go Arrays (used live by NICO) ---------------------------------
// ARRAYS: fixed-size (3 puroks, size never changes).
// index 0 = Purok 1, index 1 = Purok 2, index 2 = Purok 3
var supply = [3]int{InitialSupplyEach, InitialSupplyEach, InitialSupplyEach}

// JAVA COMPARISON: int[] supply = {3, 3, 3};

// Copying an array in Go copies ALL the data (value type), not a
// reference. NICO demonstrates this live further down.
var originalSupply = supply

// JAVA COMPARISON:
// int[] originalSupply = supply; // this copies the REFERENCE, not the data

// ================================================================
// ZYMON's SECTION -- Go Functions
// ================================================================

// FUNCTIONS + SWITCH: map a purok number to its zone name.
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

// FUNCTIONS + tagless SWITCH: classify by age.
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

// FUNCTIONS + CONDITIONS + OPERATORS: check if a resident currently
// qualifies for ayuda.
func isEligible(r Resident) bool {
	hasSupply := supply[r.Purok-1] > 0 // comparison operator: >
	isAdult := r.Age >= MinAge         // comparison operator: >=
	notYetClaimed := !r.HasClaimed     // logical operator: !
	return isAdult && hasSupply && notYetClaimed
}

// FUNCTIONS: process one resident's claim.
// Takes a POINTER (*Resident) so the change actually sticks on the
// real entry inside the slice, not just a local copy.
func processClaim(r *Resident) bool {
	if !isEligible(*r) {
		return false
	}
	r.HasClaimed = true
	supply[r.Purok-1]-- // arithmetic operator: decrement
	return true
}

// FUNCTIONS: explains why a resident was rejected.
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

// A function that RETURNS MULTIPLE VALUES -- a Go feature Java does
// not have without wrapping values in an object first.
func supplyStatus(purok int) (remaining int, isEmpty bool) {
	remaining = supply[purok-1]
	isEmpty = remaining <= 0
	return remaining, isEmpty
}

// ================================================================
// JENNY -- GO SYNTAX & GO COMMENTS
// ================================================================
func demoJenny() {
	fmt.Println("\n========================================")
	fmt.Println(" JENNY: GO SYNTAX & GO COMMENTS")
	fmt.Println("========================================")

	// This whole file already demonstrates the basic building
	// blocks of Go syntax:
	//   1. "package main"   -> every Go file starts by declaring
	//                          which package it belongs to.
	//                          "main" is special: it marks this as
	//                          a runnable program, not a library.
	//   2. "import \"fmt\""  -> brings in Go's standard library
	//                          package for formatted I/O.
	//   3. "func main() {}" -> the entry point. Execution always
	//                          starts here, just like public static
	//                          void main(String[] args) in Java.
	//   4. Curly braces "{ }" define blocks, semicolons are NOT
	//                          typed by the programmer (Go inserts
	//                          them automatically at compile time).

	/* This is a BLOCK COMMENT.
	   It can span multiple lines, just like a block comment in Java.
	   Used above for the file header and struct/function notes. */

	// This is a LINE COMMENT (used everywhere in this file to
	// label which language feature each block demonstrates).

	fmt.Println("A Go source file = package declaration + imports + code.")
	fmt.Println("Comments: // for single line, /* ... */ for blocks.")
}

// ================================================================
// LUIS -- GO VARIABLES & GO CONSTANTS
// ================================================================
func demoLuis() {
	fmt.Println("\n========================================")
	fmt.Println(" LUIS: GO VARIABLES & GO CONSTANTS")
	fmt.Println("========================================")

	// ---- Declare Variables ----
	var residentCount int // declared with an explicit type, zero-valued (0) until assigned
	residentCount = 11
	fmt.Println("Explicit var declaration:", residentCount)

	// Short variable declaration: Go infers the type. Not possible
	// in Java, where you must always write the type.
	barangayName := "Baguio City"
	fmt.Println("Short declaration (:=):", barangayName)

	// ---- Declare Multiple Variables ----
	var purokA, purokB, purokC = 1, 2, 3 // multiple vars, one line
	fmt.Println("Multiple var declare:", purokA, purokB, purokC)

	var (
		staffName string = "Kap. Ramel"
		staffAge  int    = 45
	)
	fmt.Println("Grouped var block:", staffName, staffAge)

	// ---- Naming Rules ----
	// - Must start with a letter or underscore, then letters/digits/underscore.
	// - Case matters, AND case controls visibility outside the package:
	//     lowercase first letter  -> unexported (private to the package)
	//     Uppercase first letter  -> exported   (visible to other packages)
	localHelper := "unexported: only visible inside this package"
	// Example already in this file: `Resident` (exported struct/type),
	// `purokNames` (unexported package variable).
	fmt.Println(localHelper)

	// ---- Go Constants ----
	// MinAge and InitialSupplyEach (declared near the top of this file)
	// are constants: fixed for the entire run of the program.
	fmt.Println("Constants in use -> MinAge:", MinAge, "| InitialSupplyEach:", InitialSupplyEach)
}

// ================================================================
// LEX -- GO OUTPUT & GO DATA TYPES
// ================================================================
func demoZymon() {
	fmt.Println("\n========================================")
	fmt.Println(" LEX: GO OUTPUT & GO DATA TYPES")
	fmt.Println("========================================")

	// ---- Output Functions ----
	fmt.Print("fmt.Print does not add a newline. ")
	fmt.Println("fmt.Println adds one automatically.")
	fmt.Printf("fmt.Printf formats output like this: %s is %d years old.\n", "Ramel", 45)

	// ---- Formatting Verbs (printf) ----
	name := "Kisha"
	age := 25
	height := 1.62
	isResident := true

	fmt.Printf("%%s (string):  %s\n", name)
	fmt.Printf("%%d (int):     %d\n", age)
	fmt.Printf("%%.2f (float, 2 decimals): %.2f\n", height)
	fmt.Printf("%%t (bool):    %t\n", isResident)
	fmt.Printf("%%v (any value, default format): %v\n", Resident{Name: "Kisha", Age: 25, Purok: 2})
	fmt.Printf("%%T (the type itself): %T\n", height)

	// ---- Go Data Types ----
	var wholeNumber int = 30         // integer
	var decimalNumber float64 = 12.5 // floating point
	var text string = "ayuda"        // string
	var flag bool = false            // boolean
	fmt.Println("\nBasic data types:")
	fmt.Println(" int:", wholeNumber, "| float64:", decimalNumber, "| string:", text, "| bool:", flag)
}

// ================================================================
// NICO -- GO ARRAYS & GO SLICES
// ================================================================
func demoNico() []Resident {
	fmt.Println("\n========================================")
	fmt.Println(" NICO: GO ARRAYS & GO SLICES")
	fmt.Println("========================================")

	// ---- Go Arrays ----
	// `supply` (declared at the top of the file) is a fixed-size
	// array: [3]int. Its size is part of its TYPE -- a [3]int and a
	// [5]int are different types to the Go compiler, not just two
	// arrays of different length.
	fmt.Println("supply array (fixed size, 3 puroks):", supply)

	// Arrays behave like VALUES, not references. Copying one
	// duplicates every element into new memory -- the opposite of
	// Java, where array assignment copies a reference.
	fmt.Println("originalSupply is a full, independent copy:", originalSupply)

	// ---- Go Slices: Create Slice ----
	// SLICES: dynamic-size, built on top of an array under the hood.
	residents := []Resident{
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

	// JAVA COMPARISON:
	// ArrayList<Resident> residents = new ArrayList<>(List.of(
	//     new Resident("Ramel", 45, 1, false),
	//     new Resident("Lex", 16, 1, false)
	// ));

	fmt.Printf("residents slice: len=%d, cap=%d\n", len(residents), cap(residents))

	// ---- Go Slices: Modify Slice (grow with make + append) ----
	walkIns := make([]Resident, 0, 2) // length 0, capacity 2
	walkIns = append(walkIns, Resident{Name: "Arnold", Age: 38, Purok: 2})
	walkIns = append(walkIns, Resident{Name: "Tunying", Age: 86, Purok: 1})

	// JAVA COMPARISON: ArrayList<Resident> walkIns = new ArrayList<>(2);
	// (Java's constructor only reserves capacity; it never lets you
	// inspect that reserved capacity afterward the way Go's cap() does.)

	residents = append(residents, walkIns...)
	fmt.Printf("residents after walk-ins appended: len=%d, cap=%d\n", len(residents), cap(residents))

	// Slicing an existing slice does NOT copy data -- it's a window
	// into the SAME underlying array.
	purok1Residents := residents[0:6]
	fmt.Println("Purok 1 residents (via slicing):")
	for _, r := range purok1Residents {
		fmt.Println(" -", r.Name)
	}

	purok1Residents[0].HasClaimed = true // JAVA COMPARISON: purok1Residents.get(0).setHasClaimed(true);
	fmt.Printf("Shared memory check -> residents[0] (%s) HasClaimed = %v (changed via purok1Residents!)\n",
		residents[0].Name, residents[0].HasClaimed)
	purok1Residents[0].HasClaimed = false // reset before the real simulation runs

	return residents
}

// ================================================================
// NATES -- GO OPERATORS, GO CONDITIONS, GO SWITCH, GO LOOPS
// ================================================================
func demoNates() {
	fmt.Println("\n========================================")
	fmt.Println(" NATES: OPERATORS, CONDITIONS, SWITCH, LOOPS")
	fmt.Println("========================================")

	// ---- Go Operators ----
	a, b := 7, 2
	fmt.Println("Arithmetic:  a+b =", a+b, "| a-b =", a-b, "| a*b =", a*b, "| a/b =", a/b, "| a%b =", a%b)

	x := 10
	x += 5 // assignment operator
	fmt.Println("Assignment:  x += 5 ->", x)

	fmt.Println("Comparison:  a > b ->", a > b, "| a == b ->", a == b)
	fmt.Println("Logical:     (a > b) && (x > 0) ->", (a > b) && (x > 0))

	// Go has no ternary operator (a > b ? a : b like Java) -- an
	// if-else is required instead.
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	fmt.Println("No ternary operator in Go -- max via if/else:", max)

	// ---- Go Conditions ----
	sampleAge := 16
	if sampleAge >= 18 {
		fmt.Println("Condition: Adult")
	} else if sampleAge >= 13 {
		fmt.Println("Condition: Teen")
	} else {
		fmt.Println("Condition: Kid")
	}

	// ---- Go Switch: Single-case with fallthrough ----
	fmt.Println("\nSwitch (single-case, with fallthrough):")
	purok := 1
	switch purok {
	case 1:
		fmt.Println(" ", purokZone(1))
		fallthrough // Go does NOT fall through by default -- must opt in
	case 2:
		fmt.Println(" ", purokZone(2))
	case 3:
		fmt.Println(" ", purokZone(3))
	default:
		fmt.Println("  Unknown Zone")
	}

	// ---- Go Switch: Multi-case ----
	fmt.Println("\nSwitch (multi-case, comma-separated):")
	score := 8
	switch score {
	case 1, 2, 3, 4, 5, 6:
		fmt.Println("  High level")
	case 7, 8:
		fmt.Println("  Mid level")
	case 9, 10:
		fmt.Println("  Lower level")
	default:
		fmt.Println("  Unknown level")
	}

	// ---- Go Switch: tagless (if-else chain style) ----
	fmt.Println("\nSwitch (tagless, acts like if-else):")
	fmt.Println(" ", ageGroup(70))

	// ---- Go Loops ----
	// For is the ONLY loop keyword in Go -- no separate while/do-while.

	fmt.Println("\nClassic for loop:")
	for i := 0; i < 3; i++ {
		fmt.Println(" i =", i)
	}

	fmt.Println("\nWhile-style for loop (condition only):")
	n := 0
	for n < 3 {
		fmt.Println(" n =", n)
		n++
	}

	fmt.Println("\nDo-while-style for loop (runs at least once, breaks manually):")
	m := 0
	for {
		fmt.Println(" m =", m)
		m++
		if m >= 3 {
			break
		}
	}

	fmt.Println("\nFor-each (range) with index and value:")
	zones := []string{"Zone A", "Zone B", "Zone C"}
	for i, z := range zones {
		fmt.Println(" index", i, "->", z)
	}

	fmt.Println("\nFor-each (range) ignoring the index with _:")
	for _, z := range zones {
		fmt.Println(" ", z)
	}
}

// ================================================================
// ZYMON -- GO FUNCTIONS, GO STRUCT, GO MAPS (live demonstration)
// ================================================================
// The Resident struct, purokNames map, and all the functions used
// below (isEligible, processClaim, ageGroup, purokZone,
// rejectionReason, supplyStatus) are declared near the top of this
// file -- see the "ZYMON's SECTION" banners above.
func demoLex(residents []Resident) {
	fmt.Println("\n========================================")
	fmt.Println(" ZYMON: FUNCTIONS, STRUCT & MAPS IN ACTION")
	fmt.Println("========================================")

	// STRUCT in use: each entry in `residents` is a Resident value.
	fmt.Println("Sample struct value:", residents[0])

	// MAP in use: look up a purok's display name by its number.
	fmt.Println("Map lookup purokNames[2]:", purokNames[2])
	if displayName, found := purokNames[4]; !found {
		fmt.Println("Map lookup purokNames[4]: not found (missing key) ->", displayName == "")
	}

	// Multi-return function in use.
	remaining, empty := supplyStatus(1)
	fmt.Println("supplyStatus(1) -> remaining:", remaining, "| empty:", empty)

	var totalClaimed, totalRejected int

	fmt.Println("\n--- Processing Claims ---")
	for i := 0; i < len(residents); i++ {
		r := &residents[i] // pointer so processClaim can mutate the real entry
		group := ageGroup(r.Age)
		zone := purokZone(r.Purok)

		if processClaim(r) {
			totalClaimed++
			fmt.Printf("[CLAIMED]  %-12s | %-14s | %s (%s) | %d packs left\n",
				r.Name, group, purokNames[r.Purok], zone, supply[r.Purok-1])
		} else {
			totalRejected++
			reason := rejectionReason(*r)
			fmt.Printf("[REJECTED] %-12s | %-14s | %s (%s) | reason: %s\n",
				r.Name, group, purokNames[r.Purok], zone, reason)
		}
	}

	fmt.Println("\n--- Summary Report ---")
	fmt.Println("Total Residents:", len(residents))
	fmt.Println("Total Claimed:  ", totalClaimed)
	fmt.Println("Total Rejected: ", totalRejected)

	var totalAge float64
	for _, r := range residents {
		totalAge += float64(r.Age) // type conversion: int -> float64
	}
	averageAge := totalAge / float64(len(residents))
	fmt.Printf("Average Age of Residents: %.2f\n", averageAge)

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

// ================================================================
// main() -- runs every section in presentation order.
// ================================================================
func main() {
	fmt.Println("=== Barangay Ayuda Distribution System ===")
	fmt.Println("Group Ewan | Golang | Baguio City barangay problem")

	demoJenny()
	demoLuis()
	demoZymon()
	residents := demoNico()
	demoNates()
	demoLex(residents)
}
