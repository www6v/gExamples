package reflect

import (
	"reflect"
	"strings"
	"testing"
)

// https://blog.csdn.net/m0_53157173/article/details/126346418
type i1 interface {
	F1() string
}
type s1 struct {
	h    string
	Name string
	Type int `json:"type" id:"100"`
}

func (h s1) F11() string {
	return h.Name
}

func (h s1) f21() string {
	return h.Name
}

func (h *s1) F31() string {
	return h.Name
}

func TestStructReflect(t *testing.T) {
	h := reflect.TypeOf(s1{})
	hptr := reflect.TypeOf(&s1{})
	i1 := reflect.TypeOf((*i1)(nil)).Elem()

	println(h.Align())
	println(h.FieldAlign())
	println(h.NumField())
	println(h.NumMethod())
	println(h.AssignableTo(i1))
	println(h.Implements(i1))
	println(h.ConvertibleTo(i1))

	println(hptr.AssignableTo(i1))
	println(hptr.Implements(i1))
	println(hptr.ConvertibleTo(i1))

	for i := 0; i < h.NumField(); i++ {
		f := h.Field(i)
		println("---------")
		println(f.Name)
		println(f.Type.FieldAlign())
		println(f.Type.String())
		println(f.Index)
		println(f.Anonymous)
		println(f.Offset)
		println(f.Tag)
		println("---------")
	}

	for i := 0; i < h.NumMethod(); i++ {
		f := h.Method(i)
		println("---------")
		println(f.Name)
		println(f.Type.String())
		println(f.Index)
		println(f.Func.String())
		println("---------")
	}

	m1, b := h.MethodByName("f2")
	println(b, m1.Name)

	s1, b := h.FieldByName("h")
	println(b, s1.Name)

	s2, b := h.FieldByNameFunc(func(s2 string) bool {
		// 必须唯一匹配
		return strings.Contains(s2, "e")
	})
	println(b, s2.Name)

}
