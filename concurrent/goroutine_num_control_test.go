package concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/////////////
// var wg = sync.WaitGroup{}

/// https://eddycjy.gitbook.io/golang/di-1-ke-za-tan/control-goroutine

func TestMain1(t *testing.T) {
	userCount := 10
	ch := make(chan bool, 2)
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go Read(ch, i)
	}

	wg.Wait()
}

func Read(ch chan bool, i int) {
	defer wg.Done()

	ch <- true
	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
	<-ch
}

var wg sync.WaitGroup

func TestMain2(t *testing.T) {
	userCount := 10
	ch := make(chan int, 5)

	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for d := range ch {
				fmt.Printf("go func: %d, time: %d\n", d, time.Now().Unix())
				time.Sleep(time.Second * time.Duration(d))
			}
		}()
	}

	for i := 0; i < 10; i++ {
		ch <- 1
		ch <- 2
		//time.Sleep(time.Second)
	}

	close(ch)
	wg.Wait()
}
