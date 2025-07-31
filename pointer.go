package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial value of i:", i)
	zeroval(i)

	fmt.Println("zeroval func:", i)

	zeroptr(&i)
	fmt.Println("zeroptr func:", i)

	fmt.Println("Pointer to i:", &i)
}
