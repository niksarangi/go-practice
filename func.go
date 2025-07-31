package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b

}

func triplePlus(a int, b int, c int) int {
	return a + b + c
}

func printFuncs() {
	var a, b, c int
	fmt.Print("enter a:")
	fmt.Scanf("%d\n", &a)

	fmt.Print("enter b:")

	fmt.Scanf("%d\n", &b)

	fmt.Println("summation", plus(a, b))

	fmt.Print("enter c:")
	fmt.Scanf("%d\n", &c)

	fmt.Println("summation of 3 numbers", triplePlus(a, b, c))

}
