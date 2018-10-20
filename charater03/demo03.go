package main

import "fmt"

func main() {
	var arr [2]string
	arr2 := [...]string{"hello", "keep"}
	arr3 := [5]int{1: 3, 4: 2, 3: 1}
	arr4 := [5]int{1, 2, 3, 4, 5}

	arr[0] = "hello"
	arr[1] = "array"

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	var slice []int = []int{1, 2, 3, 4, 5, 6, 7}
	slice = append(slice, 8)
	fmt.Println(slice, "\t", len(slice), "\t", cap(slice))
	fmt.Println(slice[:4], "\t", len(slice), "\t", cap(slice))
	fmt.Println(slice[4:], "\t", len(slice), "\t", cap(slice))

	var slice2 []int = make([]int, 5)
	fmt.Println(slice2, "\t", len(slice2), "\t", cap(slice2))
	var slice3 []int = slice2[:3]
	slice3 = append(slice3, 1, 2, 3, 4, 6)
	fmt.Println(slice3, "\t", len(slice3), "\t", cap(slice3))

	fmt.Println("----------------------------------------")
	for k, v := range slice3 {
		fmt.Printf("%v : %v \n", k, v)
	}

}
