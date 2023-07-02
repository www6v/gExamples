package basic

import "testing"

type MyInt int64

/// golang 类型和别名不支持隐式类型转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 2
	aPtr := &a
	// aPtr = aPtr + 1 /// 不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string /// 字符串初始化为值类型， 默认值是空字符串
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Logf("string empty")
	}
}
