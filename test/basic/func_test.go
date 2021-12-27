package basic

import (
	"fmt"
	"testing"
	"time"
)

/// 函数多值返回
func returnMultValues() (int, int) {
	return 10, 20
}

/// 函数式一等公民
func timeSpent(inner func(op int) int) func(op int) int {
	/// Decorator 包装器模式
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Print("time spend:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

/// 测试入口
func TestFn(t *testing.T) {
	a, _ := returnMultValues()
	t.Log(a)
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}
