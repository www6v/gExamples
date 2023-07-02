package basic

import (
	"fmt"
	"testing"
)

// ////////// https://zhuanlan.zhihu.com/p/522568859
///////////// 指针类型接受者

type Person1 struct {
	name string
	age  int
}

func (p Person1) say() {
	fmt.Printf("I (%p)  am %s, %d years old\n", &p, p.name, p.age)
}

func (p *Person1) older() { // 指针接受者，传递函数内部的是原类型值（指针）， 函数内的操作会体现到原指针指向的空间
	p.age = p.age + 1
	fmt.Printf("I (%p)  am %s, %d years old\n", p, p.name, p.age)
}

func TestReciever1(t *testing.T) {
	p1 := Person1{"zhangsan", 20}
	p1.older() // 虽然定义的是指针接受者，但是值类型依旧可以使用，但是会隐式传入指针值
	p1.say()
	fmt.Printf("I (%p)  am %s, %d years old\n", &p1, p1.name, p1.age)

	p2 := &Person1{"sili", 20}
	p2.older()
	p2.say()
	fmt.Printf("I (%p)  am %s, %d years old\n", p2, p2.name, p2.age)
}
