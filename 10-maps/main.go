package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)

	v1 := m["k1"]
	v3 := m["k3"]
	fmt.Println("k1:", v1, "k3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println(m, "len:", len(m))

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"k1": 1, "k2": 2}
	fmt.Println(n)

	n2 := map[string]int{"k1": 1, "k2": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	//c := n
	c := maps.Clone(n)
	n["k3"] = 3
	fmt.Println("n:", n, "c:", c)
}
