package basic

import "testing"

/// 函数是一等公民, return 函数类型
func setup(task string) func() {
	println("do some setup stuff for", task)
	return func() { println("do some teardown stuff for", task) }
}

func TestFuction(t *testing.T) {
	teardown := setup("demo")
	defer teardown()
	println("do some bussiness stuff")
}
