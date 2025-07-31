package main

import "fmt"

func printMap() {
	mapVeggie := inputVeggies()

	clear(mapVeggie)

}
func clear(m map[string]int) {
	for k := range m {
		delete(m, k)
		fmt.Println("deleted", k)

	}

}

func inputVeggies() map[string]int {
	var n int
	fmt.Print("Enter number of vegetables: ")
	fmt.Scanln(&n)
	veggies := make(map[string]int)
	for i := 0; i < n; i++ {
		var name string
		var quantity int
		fmt.Printf("Enter name of vegetable #%d: ", i+1)
		fmt.Scanln(&name)
		fmt.Printf("Enter quantity for %s: ", name)
		fmt.Scanln(&quantity)
		veggies[name] = quantity
	}
	return veggies
}
