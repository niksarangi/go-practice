package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	// Two-pointer technique to reverse
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Convert back to string
	return string(runes)
}

func main() {
	var str string
	fmt.Println("Enter string:")
	fmt.Scan(&str)
	fmt.Println("Reversed string is:", reverseString(str))
}
