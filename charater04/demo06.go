/**
Go 程序使用 error 值来表示错误状态。

与 fmt.Stringer 类似，error 类型是一个内建接口：

type error interface {
    Error() string
}
（与 fmt.Stringer 类似，fmt 包在打印值时也会满足 error。）

通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理。

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)

error 为 nil 时表示成功；非 nil 的 error 表示失败。
*/
package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v , %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "It didn't work"}
}

/*
练习：错误
从之前的练习中复制 Sqrt 函数，修改它使其返回 error 值。

Sqrt 接受到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。

创建一个新的类型

type ErrNegativeSqrt float64
并为其实现

func (e ErrNegativeSqrt) Error() string
方法使其拥有 error 值，通过 ErrNegativeSqrt(-2).Error() 调用该方法应返回 "cannot Sqrt negative number: -2"。

*注意：* 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。这是为什么呢？

修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值。
*/

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, &MyError{time.Now(), fmt.Sprintf("%v is less than zero", f)}
	} else {
		return math.Sqrt(f), nil
	}

}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	f := -20.0
	result, err := Sqrt(f)
	if err == nil {
		fmt.Println(f, "sqrt is ", result)
	} else {
		fmt.Println(err)
	}
}
