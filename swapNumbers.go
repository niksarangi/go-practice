package main

import "fmt"

func swapNumbers(a, b int) (int, int) {
	return b, a
}

func main() {
	var x, y int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	x, y = swapNumbers(x, y)

	fmt.Println("After swapping:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
}
