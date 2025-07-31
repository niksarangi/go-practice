package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func printREcursion() {
	var n int
	fmt.Println("Enter number to find its factorial")
	fmt.Scanln(&n)
	fmt.Println("Factorial is: ", fact(n))

	var sum func(n int) int
	fmt.Println("enter to find summation of numbers from 1 to n")

	sum = func(n int) int {
		if n == 1 {
			return 1
		}
		return n + sum(n-1)
	}

	fmt.Scanln(&n)
	fmt.Println("Summation is: ", sum(n))

}
