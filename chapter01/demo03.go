package main

import (
	"fmt"
	"math/rand"
	"time"
)

var golang, java, php, python int
var hello = "hello"

func main() {
	rand.Seed(time.Now().Unix())
	golang = rand.Intn(10)
	java = rand.Intn(10)
	php = rand.Intn(10)
	python = rand.Intn(10)

	fmt.Println(hello)
	fmt.Println(golang, java, php, python)
}
