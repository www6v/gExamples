package basic

import (
	"fmt"
	// gotools "gDemo/copyStruct"
)

type Src struct {
	A int
	B string
	c byte
}

type Dst struct {
	A int
	B string //byte
	c byte
}

func ArrayTest() {
	str1 := []string{"1", "2", "3", "4"}
	for index, value := range str1 {
		fmt.Println(index, ":", value)
	}
	numbers := [6]int{1, 2, 3, 5}
	for index, value := range numbers {
		fmt.Printf("第 %d 位的值 = %d\n", index, value)
	}
	p("rocky")
}

// func CopyTest() {
// 	src := &Src{3, "hello", '2'}
// 	var dst Dst
// 	gotools.StructCopy(&dst, src)
// 	fmt.Println(dst)
// }

func p(name string) {
	fmt.Println("Hello," + name)
}

// / https://studygolang.com/articles/10751
func MakeTest() {
	a := make([]int, 10, 20)
	fmt.Printf("%d, %d\n", len(a), cap(a))
	fmt.Println(a)
	b := a[:cap(a)]
	fmt.Println(b)
}
