package main

import "fmt"

func vals() (int, string) {
	return 10, "Hello, world!"
}

func main() {
	a, b := vals()
	fmt.Println("a:", a, "b:", b)

	_, c := vals()
	fmt.Println(c)
}
