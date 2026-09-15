package main

import "fmt"

/* ================================================================
   ZYMON — GO OUTPUT  (Barangay Ayuda)
   How Go prints text: Print, Println, and Printf.

   Run:  go run ./Zymon/Output/main.go
   ================================================================ */

type Resident struct {
	Name       string
	Age        int
	Purok      int
	HasClaimed bool
}

func main() {
	fmt.Println("=== Barangay Ayuda — GO OUTPUT ===")

	// ---- Print vs Println vs Printf ----
	// Print   = no extra newline at the end
	// Println = adds a newline automatically
	// Printf  = format string + verbs (like Java's System.out.printf)
	fmt.Print("fmt.Print stays on one line. ")
	fmt.Println("fmt.Println then jumps to the next line.")
	fmt.Printf("fmt.Printf inserts values: %s is %d years old.\n", "Zymon", 38)

	zymon := Resident{Name: "Zymon", Age: 38, Purok: 1, HasClaimed: false}
	packsLeft := 3
	averageAge := 36.54
	isResident := true

	fmt.Println("\n--- Formatting verbs (Printf) ---")
	fmt.Printf("%%s  string:     %s\n", zymon.Name)
	fmt.Printf("%%d  int:        %d\n", zymon.Age)
	fmt.Printf("%%.2f float:     %.2f\n", averageAge)
	fmt.Printf("%%t  bool:       %t\n", isResident)
	fmt.Printf("%%v  any value:  %v\n", zymon)
	fmt.Printf("%%T  the type:   %T\n", averageAge)

	fmt.Println("\n--- Ayuda line (same style as the main program) ---")
	fmt.Printf("[CLAIMED]  %-12s | age %d | Purok %d | %d packs left\n",
		zymon.Name, zymon.Age, zymon.Purok, packsLeft)
}
