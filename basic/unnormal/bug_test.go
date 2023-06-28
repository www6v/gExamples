package basic

import (
	"fmt"
	"testing"
	"time"
)

// / 《25 直播：如何写出优雅的 Go 代码-更多课程【ubkz.com】_一手IT课程资源+微信2268731[2]》
func TestClosure(t *testing.T) {
	data := []string{"1", "2", "3"}

	////// bugs
	// for _, v := range data {
	// 	go func() {
	// 		fmt.Println(v)
	// 	}()
	// }

	////// fixed
	for _, v1 := range data {
		go func(v1 string) {
			fmt.Println(v1)
		}(v1)
	}

	time.Sleep(3 * time.Second)
}

type person struct {
	age  int
	name string
}

func TestRange(t *testing.T) {
	// /////  bugs
	var m = map[int]person{
		1: person{11, "abc"},
		2: person{22, "def"},
	}

	var mm = map[int]*person{}
	for k, v := range m {
		mm[k] = &v
	}

	fmt.Println(mm[1], mm[2])
}
