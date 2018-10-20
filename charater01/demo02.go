package main

import(
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func swap(x, y string) (a,b string) {
	return y , x 
}

func main() {
	x := 10
	y := 20
	a := "hello"
	b := "hi"

	fmt.Println(a,":",b)
	a , b = swap(a,b)
	fmt.Println(a,":",b)
	fmt.Println(add(x , y))
	fmt.Println(sub(x,y))
}