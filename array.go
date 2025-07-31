package main

import "fmt"

func printArray() {
	i := 1
	for i <= 3 {
		fmt.Print(i)
		i = i + 1
	}

	for i := range 4 {
		fmt.Print(i)
	}

	for {
		fmt.Print(" loop")
		break
	}

	for i := 0; i < 0; i++ {
		fmt.Print(i)
	}

}
