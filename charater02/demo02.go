package main

import (
	"fmt"
	"math"
)

//牛顿法开方
func main() {
	n := 9321749.00
	result := Sqrt(n, 1)
	fmt.Printf("Sqrt %v is   %.63f \n", n, result)
	fmt.Printf("Math Lib is %.63f \n", math.Sqrt(n))
	jqd := math.Abs(n - result*result)
	libJqd := math.Abs(n - math.Sqrt(n)*math.Sqrt(n))
	fmt.Printf("JQD = %.63f \n", jqd)
	fmt.Printf("Lib JQD = %.63f \n", libJqd)

	fmt.Println("----------------------------------")
	fmt.Printf("result = %.63f \n", jqd-libJqd)
}

//递归法
func Sqrt(x float64, z float64) float64 {
	prep := z
	z -= (z*z - x) / (2 * z)
	fmt.Printf("%.60f \n", z)
	if math.Abs(prep-z) > 1e-10 {
		z = Sqrt(x, z)
	}
	return z
}
