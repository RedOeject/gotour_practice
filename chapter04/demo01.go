/**
方法就是一类带特殊的 接收者 参数的函数。
方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
记住：方法只是个带接收者参数的函数。

带指针参数的函数必须接受一个指针：
var v Vertex
ScaleFunc(v, 5)  // 编译错误！
ScaleFunc(&v, 5) // OK
而以指针为接收者的方法被调用时，接收者既能为值又能为指针：

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
由于 Scale 方法有一个指针接收者，为方便起见，Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)。


选择值或指针作为接收者
使用指针接收者的原因有二：
首先，方法能够修改其接收者指向的值。
其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Float float64

//Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/**
指针接收者
你可以为指针接收者声明方法。
这意味着对于某类型 T，接收者的类型可以用 *T 的文法。（此外，T 不能是像 *int 这样的指针。）
例如，这里为 *Vertex 定义了 Scale 方法。
指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
试着移除函数声明中的 *，观察此程序的行为如何变化。
若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。（对于函数的其它参数也是如此。）Scale 方法必须用指针接受者来更改 main 函数中声明的 Vertex 的值。
*/
func (v *Vertex) Scale(f float64) {
	v.Y *= f
	v.X *= f
}

/**
你也可以为非结构体类型声明方法。
你只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。
就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。）
*/
func (f Float) Abs() float64 {
	if f >= 0 {
		return float64(f)
	} else {
		return float64(-f)
	}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := Float(-2.0)
	fmt.Println(f.Abs())

	v2 := Vertex{2, 4}
	v2.Scale(10)
	fmt.Println(v2)
}
