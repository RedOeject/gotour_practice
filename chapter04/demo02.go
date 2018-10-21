/**
接口类型 是由一组方法签名定义的集合。
接口类型的变量可以保存任何实现了这些方法的值。

类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。
隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。
因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。
*/
package main

import (
	"fmt"
	"math"
)

type MyFloat float64
type MyVertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

//表示类型MyFloat实现了接口Abser 但无需显式声明
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *MyVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-3.14)
	v := MyVertex{-5, -4}

	//因为MyFloat和*MyVertex实现了Abser接口 所以a可以保存值
	a = f
	a = &v

	//因为是*MyVertex实现Abser接口不是MyVertex，所以不能使用a=v
	//a = v
	fmt.Printf("%T : %v \n", a, a)
	fmt.Println(a.Abs())
}
