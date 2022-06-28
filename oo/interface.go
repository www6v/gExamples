package oo

/**
如果新接口类型嵌入了多个接口类型，这些嵌入的接口类型的方法集合不能有交
集，同时嵌入的接口类型的方法集合中的方法名字，也不能与新接口中的其他方法同名。
比如我们用 Go 1.12.7 版本运行下面例子，Go 编译器就会报错
*/
type Interface1 interface {
	M1()
}

type Interface2 interface {
	M1()
	M2()
}

type Interface3 interface {
	Interface1
	Interface2 // Error: duplicate method M1
}

type Interface4 interface {
	Interface2
	M2() // Error: duplicate method M2
}
