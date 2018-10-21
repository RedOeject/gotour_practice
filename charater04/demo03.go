/**
接口值
在内部，接口值可以看做包含值和具体类型的元组：
(value, type)
接口值保存了一个具体底层类型的具体值。
接口值调用方法时会执行其底层类型的同名方法。


*/
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("< nil >")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func description(i I) {
	fmt.Printf("%T : %v \n", i, i)
}

func descriptionNone(l interface{}) {
	fmt.Printf("%T : %v \n", l, l)
}

func main() {
	var i I
	i = &T{"Hello"}
	description(i)
	i.M()

	fmt.Println("---------------------------")

	i = F(math.Pi)
	description(i)
	i.M()

	/*
		即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
		在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。
		*注意：* 保存了 nil 具体值的接口其自身并不为 nil ，即 i 不为nil，t为nil
	*/
	var t *T
	i = t
	description(i)
	i.M()

	/*
		nil 接口值既不保存值也不保存具体类型。
		为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。
	*/
	var k I
	description(k)
	//k.M()

	/*
		指定了零个方法的接口值被称为 *空接口：*
		interface{}
		空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
		空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
	*/
	var l interface{}
	descriptionNone(l)
	l = 42
	descriptionNone(l)
	l = "hello"
	descriptionNone(l)

	var m I
	m = &T{"HELLO"}
	n, ok := m.(*T)
	fmt.Println(n, ok)
}
