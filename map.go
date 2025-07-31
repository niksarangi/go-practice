package main

import (
	"fmt"
	"maps"
)

func main() {
	map1 := make(map[string]int)
	map1["Nikki"] = 7
	map1["vishwa"] = 6
	map1["sam"] = 5
	map1["harsh"] = 10
	map1["fairy"] = 11

	fmt.Println("map:", map1)
	clear(map1)
	fmt.Println("map after clearing", map1)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
