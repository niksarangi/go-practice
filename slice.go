package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 20, 30}
	fmt.Println("slice", nums)
	arr := [5]int{1, 2, 3, 4, 5}
	slicevar := arr[2:5]
	fmt.Println("slicee", slicevar)

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
}
