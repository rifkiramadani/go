package main

import (
	"fmt"
	"strings"
)

func main() {
	animals := "             hello world             "
	fmt.Println(strings.TrimSpace(animals))

	// STRING SLICE
	// animals := "kucing, anjing, monyet, gajah"
	// filter := strings.Split(animals, ",")

	// fmt.Println(filter)

	// // STRING REPLACE
	// animals := "kucing, anjing , monyet, gajah"
	// fmt.Println(strings.ReplaceAll(animals, "gajah", "jerapah"))

	// STRING CONTAINS
	// animals := "kucing, anjing, monyet, jerapah"
	// fmt.Println(strings.Contains(animals, "gajah"))

	// STRING UPPERCASE & LOWERCASE
	// nameLower := "muhammad rifky ramadani"
	// nameUpper := "MUHAMMAD RIFKY RAMADANI"

	// fmt.Println(strings.ToUpper(nameLower))
	// fmt.Println(strings.ToLower(nameUpper))
}
