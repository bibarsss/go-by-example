package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 123
	return &p
}

func main() {
	fmt.Println(person{name: "Biba"})

	fmt.Println(person{name: "Biba", age: 23})

	fmt.Println(person{age: 100})

	fmt.Println(&person{name: "Bibarys", age: 23})

	fmt.Println(newPerson("Bibka"))

	a := newPerson("Bibaka")
	a.age = 32
	fmt.Println(*a)

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
}
