package main

import (
	"fmt"
	"math/rand"
	 "time"
)

func main() {
	result := 0

	for i := 1 ;i <= 10 ; i++ {
		result += i
	}

	k := 11
	for k <= 100 {
		result += k
		k++
	}

	var l int

	for {
		seed := int64(result)
		rand.Seed(seed+time.Now().Unix())
		l = rand.Intn(100)
		result += l
		fmt.Println(l)
		if kk := result;l == 5 {
			fmt.Println("kk is", kk)
			break
		}
	}

	fmt.Println("--------------------------------")
	fmt.Println(result)
}
