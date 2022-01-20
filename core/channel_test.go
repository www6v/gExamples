package core

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
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

/// << 14 | Channel：透过代码看典型的应用模式 >>  晁岳攀
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

/// 数据传递
func TestPipeline1(t *testing.T) {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}

	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	chs[0] <- struct{}{}
	select {}

}

/// 信号通知
func TestSignalNotify(t *testing.T) {
	var closing = make(chan struct{})
	var closed = make(chan struct{})

	go func() {
		// 模拟业务处理
		for {
			select {
			case <-closing:
				return
			default:
				// ....... 业务计算
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)
	// 执行退出之前的清理动作
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("清理超时，不等了")
	}
	fmt.Println("优雅退出")
}

func doCleanup(closed chan struct{}) {
	time.Sleep((time.Minute))
	close(closed)
}

/// 任务编排 -> Or-Done 模式
/**
这是有一个任务的情况，如果有多个任务，只要有任意一个任务执行完，我们就想获得这
个信号，这就是 Or-Done 模式。
比如，你发送同一个请求到多个微服务节点，只要任意一个微服务节点返回结果，就算成
功，这个时候，就可以参考下面的实现：
*/
func TestOrDone(t *testing.T) {
	start := time.Now()
	<-or(
		sig(10*time.Second),
		sig(20*time.Second),
		sig(30*time.Second),
		sig(40*time.Second),
		sig(50*time.Second),
		sig(01*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
func or(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个或者1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2: // 2个也是一种特殊情况
			select {
			case <-channels[0]:
			case <-channels[1]:
			}

		default: //超过两个，二分法递归处理
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()
	return orDone
}

/// 任务编排 ->  扇入模式
func fanInRec(chans ...<-chan interface{}) <-chan interface{} {
	switch len(chans) {
	case 0:
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		return chans[0]
	case 2:
		return mergeTwo(chans[0], chans[1])
	default:
		m := len(chans) / 2
		return mergeTwo(
			fanInRec(chans[:m]...),
			fanInRec(chans[m:]...))
	}
}
func mergeTwo(a, b <-chan interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for a != nil || b != nil { //只要还有可读的chan
			select {
			case v, ok := <-a:
				if !ok { // a 已关闭，设置为nil
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok { // b 已关闭，设置为nil
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

/// 任务编排 ->  扇出模式
func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() { //退出时关闭所有的输出chan
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()
		for v := range ch { // 从输入chan中读取数据
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				if async { //异步
					go func() {
						out[i] <- v // 放入到输出chan中,异步方式
					}()
				} else {
					out[i] <- v // 放入到输出chan中，同步方式
				}
			}
		}
	}()
}
