package core

import (
	"fmt"
	"reflect"
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

///
func TestChannelWithReflect(t *testing.T) {
	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)
	var cases = createCases(ch1, ch2)
	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() { // recv case
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else { // send case
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
	}
}

func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}
	return cases
}

type Token struct{}

func newWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch
		fmt.Println((id + 1))
		time.Sleep(time.Second)
		nextCh <- token
	}
}
func TestPipeline1(t *testing.T) {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}

	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	chs[0] <- struct{}{}
	select {}

}
