package main

type person struct {
	name   string
	age    int
	gender string
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 36
	p.gender = "F"
	return &p
}

// func main() {
// 	fmt.Println(person{name: "bob", age: 20, gender: "M"})
// 	fmt.Println(person{name: "alice", age: 20})
// 	fmt.Println(newPerson("nikki"))
// 	fmt.Println(&person{name: "Ann", age: 40})

// 	s := person{name: "mac", age: 30}

// 	fmt.Println(s.name)
// 	fmt.Println(s.age)
// 	sp := &s
// 	fmt.Println(sp.age)
// 	sp.age = 40
// 	fmt.Println(sp.age)
// }
