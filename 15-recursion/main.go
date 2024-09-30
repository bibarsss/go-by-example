package main

import (
	"fmt"
	"slices"
)

func fact(n int) int {
	if slices.Contains([]int{0, 1}, n) {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
