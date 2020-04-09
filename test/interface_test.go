package foo

import "testing"

type Programmer interface{
	WriteHelloWorld() string
}

type GoProgrammer struct {

}

///duck 类型
/// 接口为非入侵性，实现不依赖于接口定义
func (go1 *GoProgrammer) WriteHelloWorld() string{
  return "hello  111"
}

func TestClient( t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}