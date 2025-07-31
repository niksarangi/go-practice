package main

import "fmt"

// sample program to understand variadic funtions -> takes any number of arguments

func sum(nums ...int) {
	fmt.Print(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total", total)

}

func printVariadic() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{2, 3, 4}
	sum(nums...)
}
