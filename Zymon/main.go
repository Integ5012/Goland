package main

import "fmt"

/* ================================================================
   ZYMON — GO OUTPUT & GO DATA TYPES  (Barangay Ayuda)
   ================================================================ */

func main() {
	demoOutput()
	demoDataTypes()
}

func demoOutput() {
	fmt.Println("========== GO OUTPUT ==========")

	fmt.Print("fmt.Print stays on one line. ")
	fmt.Println("fmt.Println then jumps to the next line.")
	fmt.Printf("fmt.Printf inserts values: %s is %d years old.\n", "Zymon", 38)

	name := "Zymon"
	age := 38
	averageAge := 36.54
	hasClaimed := false

	fmt.Println("\n--- Formatting verbs ---")
	fmt.Printf("%%s  string:    %s\n", name)
	fmt.Printf("%%d  int:       %d\n", age)
	fmt.Printf("%%.2f float:    %.2f\n", averageAge)
	fmt.Printf("%%t  bool:      %t\n", hasClaimed)
	fmt.Printf("%%v  any value: %v\n", name)
	fmt.Printf("%%T  the type:  %T\n", averageAge)

	fmt.Printf("\n[CLAIMED]  %-12s | age %d | %d packs left\n", name, age, 2)
	fmt.Println()
}

func demoDataTypes() {
	fmt.Println("========== GO DATA TYPES ==========")

	var residentName string = "Zymon" // Java: String
	var age int = 38                  // Java: int
	var purok int = 1
	var averageAge float64 = 36.54 // Java: double
	var hasClaimed bool = false    // Java: boolean

	fmt.Println("string :", residentName)
	fmt.Println("int    :", age, "| purok:", purok)
	fmt.Println("float64:", averageAge)
	fmt.Println("bool   :", hasClaimed)

	fmt.Println("\n--- %T = the type ---")
	fmt.Printf("residentName %T | age %T | averageAge %T | hasClaimed %T\n",
		residentName, age, averageAge, hasClaimed)

	fmt.Println("\n--- Convert int -> float64 for average ---")
	totalAge := 45 + 16 + 38
	count := 3
	average := float64(totalAge) / float64(count)
	fmt.Printf("Ramel + Lex + Zymon = %d  ->  average %.2f\n", totalAge, average)

	fmt.Println("\n--- Zero values ---")
	var emptyName string
	var emptyAge int
	var emptyFlag bool
	fmt.Printf("string %q | int %d | bool %t\n", emptyName, emptyAge, emptyFlag)
}
