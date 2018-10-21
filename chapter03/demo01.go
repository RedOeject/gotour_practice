package main

import (
	"fmt"
)

func main() {
	i := 42
	var p *int
	k := &i
	fmt.Printf("%T : %v \n", p, p)
	fmt.Printf("%T : %v \n", k, k)
	fmt.Printf("%T : %v \n", *k, *k)
	*k = 43
	fmt.Printf("%T : %v \n", i, i)
	j := 2018
	p = &j
	*p = *p / j
	fmt.Printf("%T : %v \n", j, j)
}
