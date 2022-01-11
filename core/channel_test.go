package core

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			fmt.Println(v)
		}
	}
}

/// goroutine 泄漏的问题也很常见，下面的代码也是一个实际项目中的例子：
func TestChannelTest(t *testing.T) {

	startingTime := time.Now().UTC()
	time.Sleep(10 * time.Millisecond)
	endingTime := time.Now().UTC()
	var duration time.Duration = endingTime.Sub(startingTime)
	timeout := duration

	ch := make(chan bool)
	go func() {
		// 模拟处理耗时的业务
		time.Sleep((timeout + time.Second))
		ch <- true // block
		fmt.Println("exit goroutine")
	}()

	select {
	case result := <-ch:
		// return result
		fmt.Println(result)
	case <-time.After(timeout):
		// return false
		fmt.Println(false)
	}
}
