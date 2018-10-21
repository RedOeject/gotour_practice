/**
Stringer
fmt 包中定义的 Stringer 是最普遍的接口之一。

type Stringer interface {
    String() string
}

Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	name string
	age  int
}

type IPAddr [4]byte

func (p Person) String() string {
	return fmt.Sprintf("I am %v , I am %v years old.", p.name, p.age)
}

func (ipAddr IPAddr) String() string {
	var tmp string
	for _, v := range ipAddr {
		tmp += strconv.Itoa(int(v))
		tmp += "."
	}
	tmp = strings.Trim(tmp, ".")
	return tmp
}

func main() {
	a := Person{"Jun", 22}
	b := Person{"Yu", 23}

	fmt.Println(a)
	fmt.Println(b)

	/*
		练习：Stringer
		通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。

		例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。
	*/

	hosts := map[string]IPAddr{
		"googleDNS": {8, 8, 8, 8},
		"localhost": {127, 0, 0, 1},
	}

	for name, ip := range hosts {
		fmt.Println(name, ":", ip)
	}
}
