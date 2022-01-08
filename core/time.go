package core

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

func Test_Timer() {
	/*
	   用sleep实现定时器
	*/
	fmt.Println(time.Now())
	time.Sleep(time.Second)
	fmt.Println(time.Now())
	/*
	   用timer实现定时器
	*/
	timer := time.NewTimer(time.Second)
	fmt.Println(<-timer.C)
	/*
	   用after实现定时器
	*/
	fmt.Println(<-time.After(time.Second))

}
