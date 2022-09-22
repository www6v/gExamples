package concurrent

import "fmt"

// 33｜并发：小channel中蕴含大智慧 - Tony Bai
func TestChannelType() {
	ch1 := make(chan<- int, 1)
	ch2 := make(<-chan int, 1)
	// <-ch1     // invalid operation: <-ch1 (receive from send-only type chan<- in
	// ch2 <- 13 // invalid operation: ch2 <- 13 (send to receive-only type <-chan
	ch1 <- 17
	a := <-ch2

	fmt.Print(a)
}
