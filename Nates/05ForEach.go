package main

import "fmt"

// Go
func main() {
	residents := []string{"Aling Nena", "Mang Tomas", "Mang Kiko"}
	
	fmt.Println()
	for _, residents := range residents {
		fmt.Println(residents)
	}
}

//for i, residents := range residents {
//	fmt.Println(i, residents)
//}
//}

//fmt.Println()
//for _, residents := range residents {
//fmt.Println(residents)
//}
