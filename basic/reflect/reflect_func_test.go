package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// https://blog.csdn.net/m0_53157173/article/details/126346418
func TestFuncReflect(tt *testing.T) {
	h := func(p1 int, p2 ...string) string { return fmt.Sprint(p1, p2) }
	t := reflect.TypeOf(h)
	v := reflect.ValueOf(h)
	println(t.IsVariadic()) // 是否包含可变参数
	println(t.Kind().String())
	println(t.NumIn()) // 输入参数
	for i := 0; i < t.NumIn(); i++ {
		p := t.In(i) // 第i个输入参数
		println(p.Kind().String())
	}
	println(t.NumOut()) // 输出参数
	for i := 0; i < t.NumOut(); i++ {
		o := t.Out(i) // 第i个输出参数
		println(o.Kind().String())
	}

	// 参数调用
	call := v.Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf("1"), reflect.ValueOf("1")})
	for _, value := range call {
		println(value.String())
	}

	// 可变参数调用
	calls := v.CallSlice([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf([]string{"1", ""})})
	for _, value := range calls {
		println(value.String())
	}
}
