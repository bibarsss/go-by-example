package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter a: ")
	aInput, err := fmt.Scanln(&a)
	if err != nil {
		return
	}
	_ = aInput

	fmt.Print("Enter b: ")
	bInput, err := fmt.Scanln(&b)
	if err != nil {
		return
	}
	_ = bInput

	fmt.Println("Sum is:", a+b)
}
