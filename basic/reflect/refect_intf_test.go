package reflect

import (
	"reflect"
	"testing"
)

// https://blog.csdn.net/m0_53157173/article/details/126346418
type i interface {
	F1() string
	F3() string
}
type s struct {
	h    string
	Name string
	Type int `json:"type" id:"100"`
}

func (h s) F1() string {
	return h.Name
}

func (h s) f2() string {
	return h.Name
}

func (h *s) F3() string {
	return h.Name
}

func TestIntfReflect(t *testing.T) {

	t1 := reflect.TypeOf((*i)(nil)).Elem()
	s1 := reflect.TypeOf(&s{})
	s2 := reflect.TypeOf(s{})

	println(s1.Implements(t1))
	println(s2.Implements(t1))

	println(t1.Align())
	println(t1.FieldAlign())
	println(t1.NumMethod())

	for i := 0; i < t1.NumMethod(); i++ {
		f := t1.Method(i)
		println("---------")
		println(f.Name)
		println(f.Type.String())
		println(f.Index)
		println(f.Func.String())
		println("---------")
	}

	m1, b := t1.MethodByName("F1")
	println(b, m1.Name)

}
