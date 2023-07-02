package basic

import (
	"fmt"
	"testing"
)

// ////////  https://zhuanlan.zhihu.com/p/522568859
// ////////  值类型方法接收者

type Person struct {
	name string
	age  int
}

func (p Person) say() {
	fmt.Printf("I (%p) ma %s, %d years old \n", &p, p.name, p.age)
}

func (p Person) older() { // 值类型方法接受者： 接受者是原类型值的副本
	p.age = p.age + 1
	fmt.Printf("I (%p) am %s, %d years old\n", &p, p.name, p.age)
}

func TestReciever(t *testing.T) {
	p1 := Person{name: "zhangsan", age: 20}
	p1.older()
	p1.say()
	fmt.Printf("I (%p) am  %s, %d years old\n", &p1, p1.name, p1.age)

	p2 := &Person{name: "sili", age: 20}
	p2.older() // 即使定义的是值类型接受者， 指针类型依旧可以使用，但我们传递进去的还是值类型的副本
	p2.say()
	fmt.Printf("I (%p) am %s, %d years old\n", p2, p2.name, p2.age)
}
