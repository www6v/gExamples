package oo

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

type MyInt int

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}

type t struct {
	a int
	b int
}
type S struct {
	*MyInt
	t
	io.Reader
	s string
	n int
}

func TestStruct(test *testing.T) {
	m := MyInt(17)
	r := strings.NewReader("hello, go")
	s := S{
		MyInt: &m,
		t: t{
			a: 1,
			b: 2,
		},
		Reader: r,
		s:      "demo",
	}

	/// delegate “代理”
	// var sl = make([]byte, len("hello, go"))
	// s.Reader.Read(sl)
	// fmt.Println(string(sl)) // hello, go
	// s.MyInt.Add(5)
	// fmt.Println(*(s.MyInt)) // 22

	/**  “继承”
		这种“障眼法”的工作机制是这样的，当我们通过结构体类型 S 的变量 s 调用 Read 方法
	时，Go 发现结构体类型 S 自身并没有定义 Read 方法，于是 Go 会查看 S 的嵌入字段对
	应的类型是否定义了 Read 方法。这个时候，Reader 字段就被找了出来，之后 s.Read 的
	调用就被转换为 s.Reader.Read 调用。
	*/
	var sl = make([]byte, len("hello, go"))
	s.Read(sl)
	fmt.Println(string(sl))
	s.Add(5)
	fmt.Println(*(s.MyInt))
}
