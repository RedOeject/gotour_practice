package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{"hh", 1}
	point := &p
	editName(point)
	fmt.Println(point.name)
}

func editName(person *Person) {
	person.name = "nihao"
}
