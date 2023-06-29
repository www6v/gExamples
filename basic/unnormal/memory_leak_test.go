package basic

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// https://cloud.tencent.com/developer/article/2134737
// memory leak - slice
func TestSlice(t *testing.T) {
	slice1 := []int{3, 4, 5, 6, 7}
	slice2 := slice1[1:3]

	fmt.Printf("slice1 addr: %p", &slice1)
	fmt.Println()
	fmt.Printf("slice2 addr: %p", &slice2)
	fmt.Println()

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%v:[%v]  ", slice1[i], &slice1[i])
	}
	fmt.Println()
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("%v:[%v]  ", slice2[i], &slice2[i])
	}
	fmt.Println()
}

// fixed
func TestSlice1(t *testing.T) {
	initSlice := []int{3, 4, 5, 6, 7}
	//partSlice := initSlice[1:3]

	var partSlice []int
	partSlice = append(partSlice, initSlice[1:3]...) // append

	fmt.Printf("initSlice addr: %p", &initSlice)
	fmt.Println()
	fmt.Printf("partSlice addr: %p", &partSlice)
	fmt.Println()

	for i := 0; i < len(initSlice); i++ {
		fmt.Printf("%v:[%v]  ", initSlice[i], &initSlice[i])
	}
	fmt.Println()
	for i := 0; i < len(partSlice); i++ {
		fmt.Printf("%v:[%v]  ", partSlice[i], &partSlice[i])
	}
	fmt.Println()
}

// fixed
func TestSlice2(t *testing.T) {
	initSlice := []int{3, 4, 5, 6, 7}
	//partSlice := initSlice[1:3]

	partSlice := make([]int, 2)
	copy(partSlice, initSlice[1:3]) //copy

	fmt.Printf("initSlice addr: %p", &initSlice)
	fmt.Println()
	fmt.Printf("partSlice addr: %p", &partSlice)
	fmt.Println()

	for i := 0; i < len(initSlice); i++ {
		fmt.Printf("%v:[%v]  ", initSlice[i], &initSlice[i])
	}
	fmt.Println()
	for i := 0; i < len(partSlice); i++ {
		fmt.Printf("%v:[%v]  ", partSlice[i], &partSlice[i])
	}
	fmt.Println()
}

// memory leak - ticker
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop() // 这个stop一定不能漏了, 否则会导致内存泄漏
	go func(ticker *time.Ticker) {
		for {
			select {
			case value := <-ticker.C:
				fmt.Println(value)
			}
		}
	}(ticker)
	time.Sleep(time.Second * 5)
	fmt.Println("finish!!!")
}

// memory leak - goroutine
func TestSend(t *testing.T) {
	ch := make(chan int)
	fmt.Println("num of go start: ", runtime.NumGoroutine())
	time.Sleep(time.Second)

	for i := 0; i < 5; i++ { // 向channel发送5次
		go func(ii int) {
			ch <- ii
			fmt.Println("send to chan: ", ii)
		}(i)
	}

	go func() { // 只从channel接收一次
		value := <-ch
		fmt.Println("recv from chan: ", value)
	}()

	time.Sleep(time.Second)
	fmt.Println("num of go end: ", runtime.NumGoroutine())
}

func TestSend1(t *testing.T) {
	ch := make(chan int, 2)
	fmt.Println("num of go start: ", runtime.NumGoroutine())
	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		go func(ii int) {
			ch <- ii
			fmt.Println("send to chan: ", ii)
		}(i)
	}

	go func() {
		value := <-ch
		fmt.Println("recv from chan: ", value)
	}()

	time.Sleep(time.Second)
	fmt.Println("num of go end: ", runtime.NumGoroutine())
}
