package  core

import (
	"fmt"
	"time"
)

func Test_Sleep() {
	for i := 0; i < 3; i++ {
		fmt.Println("begin", time.Now().Format("2006-01-02_15:04:05"))
		fmt.Println("Do something 1s")
		time.Sleep(time.Second * 1)
		fmt.Println("end", time.Now().Format("2006-01-02_15:04:05"))
		time.Sleep(time.Second * 5)
	}
}

func Test_Tick() {
	t1 := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t1.C:
			fmt.Println("begin", time.Now().Format("2006-01-02_15:04:05"))
			fmt.Println("Do something 1s")
			time.Sleep(time.Second * 1)
			fmt.Println("end", time.Now().Format("2006-01-02_15:04:05"))
		}
	}
}
