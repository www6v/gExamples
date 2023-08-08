// 使用反射的例子
package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())
	fmt.Println("value:", reflect.ValueOf(x).Float())
	fmt.Println("value:", reflect.ValueOf(x).Interface())
	fmt.Println("value:", reflect.ValueOf(x).Interface().(float64))
}

// https://blog.csdn.net/m0_53157173/article/details/126346418
// 第一法则 从 interface{} 变量可以反射出反射对象；
func TestReflect1(t *testing.T) {
	author := "draven"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))
}

// 第二法则 从反射对象可以获取 interface{} 变量；
func TestReflect2(t *testing.T) {
	v := reflect.ValueOf(1)
	// v.Interface().(int)
	fmt.Println(v.Interface().(int))

}

// 第三法则
//
//	要修改反射对象，其值必须可设置；
//
// 从反射对象可以修改 interface{} 变量的值。
func TestReflect3(t *testing.T) {
	author := "draven"
	v := reflect.ValueOf(&author)
	v.Elem().SetString("draven1")
	fmt.Println(author)
}

func TestReflect5(t *testing.T) {
	// 这段代码会报错
	// i := 1
	// v := reflect.ValueOf(i)
	// v.SetInt(10)
	// fmt.Println(i)

	i := 1
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(10)
	fmt.Println(i)

	i1 := 1
	v1 := &i1
	*v1 = 10
	fmt.Println(i1)
}
