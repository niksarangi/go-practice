package main

import "fmt"

func values() (int, int) {
	return 3, 7
}

func printMultipleREturn() {
	a, b := values()
	fmt.Println(a)
	fmt.Println(b)

	c, _ := values()
	fmt.Println(c)

	d, _ := values()
	fmt.Println(d)

}

// package main

// import "fmt"

// func main() {
// 	var name string
// 	fmt.Print("Enter your name: ")
// 	_, _ = fmt.Scanln(&name) // both return values are ignored
// 	fmt.Println("Hello,", name)
// }
