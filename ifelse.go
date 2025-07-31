package main

import "fmt"

func printIfElse() {
	var age int
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	if age < 18 && age > 0 {
		fmt.Println("Underage")
	} else if age > 21 && age < 100 {
		fmt.Println("Legal age")
	} else {
		fmt.Println("enter valid age")
	}

}
