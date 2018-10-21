/*
Reader
io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。

Go 标准库包含了该接口的许多实现，包括文件、网络连接、压缩和加密等等。

io.Reader 接口有一个 Read 方法：

func (T) Read(b []byte) (n int, err error)
Read 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。

示例代码创建了一个 strings.Reader 并以每次 8 字节的速度读取它的输出。
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
练习：rot13Reader
有种常见的模式是一个 io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。

例如，gzip.NewReader 函数接受一个 io.Reader（已压缩的数据流）并返回一个同样实现了 io.Reader 的 *gzip.Reader（解压后的数据流）。

编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader，通过应用 rot13 代换密码对数据流进行修改。

rot13Reader 类型已经提供。实现 Read 方法以满足 io.Reader。
*/
type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b = b + 13
	case 'M' < b && b <= 'Z':
		b = b - 13
	case 'a' <= b && b <= 'm':
		b = b + 13
	case 'm' < b && b <= 'z':
		b = b - 13
	}
	return b
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, nil
}

func main() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v \n", n, err, b)
		fmt.Printf("b[:n] = %q\n \n", b[:n])
		if err == io.EOF {
			fmt.Println(err)
			break
		}
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	k := rot13Reader{s}
	io.Copy(os.Stdout, &k)
}
