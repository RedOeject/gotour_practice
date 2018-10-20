package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	rand.Seed(time.Now().Unix())
	// rand.Seed()
	fmt.Println("My favorite number is ",rand.Intn(100));
	fmt.Println(math.Pi)
}