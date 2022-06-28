package main

import (
	"fmt"
	"time"
)

func main() {
	//if ifaceType != nil {
	//
	//}
	//
	//for i := 0; i < ifaceType.Len(); i++ {
	//
	//}

	var a = make(chan struct{})
	close(a)

	ch1 := make(chan bool)
	//close(ch1)

	go func() {
		//ch1 <- true
		ch1 <- false
	}()

	select {
	case result := <-ch1:
		fmt.Println(result)
	case <-time.After(1):
		fmt.Println("a")
	}

}
