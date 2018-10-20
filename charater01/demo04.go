package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe                  = false
	MathMaxInt            = math.MaxInt64
	MaxInt     uint64     = 1<<64 - 1
	z          complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("%T : %v \n", ToBe, ToBe)
	fmt.Printf("%T : %v \n", MathMaxInt, MathMaxInt)
	fmt.Printf("%T : %v \n", MaxInt, MaxInt)
	fmt.Printf("%T : %v \n", z, z)
}
