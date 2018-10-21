package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name string
	age  int
}

var family map[string]Person

func main() {
	family = make(map[string]Person)
	family["father"] = Person{"wei", 42}
	family["mother"] = Person{"yan", 40}
	family["son"] = Person{"jun", 21}

	delete(family, "son")

	fmt.Println(family)
	elem, exist := family["father"]
	if !exist {
		fmt.Println("no exist")
	} else {
		fmt.Println(elem)
	}
	visit()
}

func visit() {
	name := family["son"].name
	names := strings.Fields(name)
	fmt.Println(names)
}
