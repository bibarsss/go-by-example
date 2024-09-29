package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func summ(nums ...int) int {
	r := 0
	for _, num := range nums {
		r += num
	}
	return r
}

func main() {
	res := plus(4, 5)
	fmt.Println("4+5 =", res)

	sum := summ(1, 2, 3, 4, 5)
	fmt.Println("sum", sum)
}
