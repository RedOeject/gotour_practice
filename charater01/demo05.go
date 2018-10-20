package main

import "fmt"

const (
	MAN = 1
	WOMAN = 0
	WHO = iota
)

func main (){
	fmt.Println(MAN, WOMAN, WHO)
}