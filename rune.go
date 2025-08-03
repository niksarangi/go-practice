package main

import (
	"fmt"
)

func countRunes(s string) int {

	count := 0
	for range s {
		count++
	}
	fmt.Println("Total number of runes in string:", count)

	return count
}

// func main() {
// 	var s string
// 	fmt.Println("Enter string:")
// 	fmt.Scanln(&s)
// 	countRunes(s)

// 	fmt.Println("String length:", len(s))

// 	for i := 0; i < len(s); i++ { //prints each byte in the string
// 		fmt.Printf("%x ", s[i])
// 	}
// 	fmt.Println()
// 	fmt.Println("Rune count:", utf8.RuneCountInString(s))

// }
