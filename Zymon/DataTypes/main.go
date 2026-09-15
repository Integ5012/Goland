package main

import "fmt"

/* ================================================================
   ZYMON — GO DATA TYPES  (Barangay Ayuda)
   Basic types used in the ayuda program: string, int, float64, bool.
   ================================================================ */

func main() {
	fmt.Println("=== Barangay Ayuda — GO DATA TYPES ===")

	// string  -> text (Java: String)
	var name string = "Zymon"

	// int     -> whole number (Java: int)
	var age int = 38
	var purok int = 1
	var packsLeft int = 3

	// float64 -> decimal number (Java: double)
	var averageAge float64 = 36.54

	// bool    -> true or false (Java: boolean)
	var hasClaimed bool = false
	var isAdult bool = age >= 18

	fmt.Println("--- Declared values ---")
	fmt.Println("string  name:       ", name)
	fmt.Println("int     age:        ", age)
	fmt.Println("int     purok:      ", purok)
	fmt.Println("int     packsLeft:  ", packsLeft)
	fmt.Println("float64 averageAge: ", averageAge)
	fmt.Println("bool    hasClaimed: ", hasClaimed)
	fmt.Println("bool    isAdult:    ", isAdult)

	fmt.Println("\n--- %T prints the type itself ---")
	fmt.Printf("name       -> %T\n", name)
	fmt.Printf("age        -> %T\n", age)
	fmt.Printf("averageAge -> %T\n", averageAge)
	fmt.Printf("hasClaimed -> %T\n", hasClaimed)

	fmt.Println("\n--- Type conversion (int -> float64) ---")
	// You cannot mix int and float64 in one expression without converting.
	totalAge := 45 + 16 + 38 // int
	count := 3               // int
	average := float64(totalAge) / float64(count)
	fmt.Printf("Ramel + Lex + Zymon ages: %d\n", totalAge)
	fmt.Printf("Average age: %.2f  (type %T)\n", average, average)

	fmt.Println("\n--- Zero values (declared, not assigned) ---")
	var emptyName string
	var emptyAge int
	var emptyFlag bool
	fmt.Printf("string zero: %q | int zero: %d | bool zero: %t\n",
		emptyName, emptyAge, emptyFlag)
}
