package foo

import (
	"time"
	"fmt"
	"testing"
)

type IntConv func(op int) int
func timeSpent1( inner IntConv) IntConv {
	/// Decorator 包装器模式
	return func(n int) int {
		start:= time.Now()
		ret:= inner(n)

		fmt.Print("time spend:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc1(op int) int {
	time.Sleep(time.Second*1)
	return op
}

/// 测试入口
func TestFn1(t *testing.T) {
	//a,_ := returnMultValues()
	//t.Log(a)
	tsSF:= timeSpent1(slowFunc1)
	t.Log( tsSF(10) )
}