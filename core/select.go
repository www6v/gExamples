package core

import (
	"fmt"
	"time"
)

func SelectTest() {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {

		time.Sleep(4 * time.Second)
		close(c)
	}()

	go func() {

		//time.Sleep(3*time.Second)
		time.Sleep(5 * time.Second)
		ch1 <- 3
	}()

	go func() {

		//time.Sleep(3*time.Second)
		time.Sleep(5 * time.Second)
		ch2 <- 5
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	case <-ch1:
		fmt.Printf("ch1 case...")
	case <-ch2:
		fmt.Printf("ch2 case...")
		//default:
		//
		//	fmt.Printf("default go...")
	}
}

var ch1 chan int
var ch2 chan int
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func SelectTest1() {

	select {
	case getChan(0) <- getNumber(2):

		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):

		fmt.Println("2th case is selected.")
	default:

		fmt.Println("default!.")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)

	return numbers[i]
}
func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)

	return chs[i]
}

func SelectTest2() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 3
	ch2 <- 5

	select {
	case <-ch1:
		fmt.Println("ch1 selected.")
		break
		fmt.Println("ch1 selected after break")
	case <-ch2:

		fmt.Println("ch2 selected.")
		fmt.Println("ch2 selected without break")
	}
}
