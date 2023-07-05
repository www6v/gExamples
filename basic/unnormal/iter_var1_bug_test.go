package basic

import (
	"fmt"
	"testing"
)

// / https://blog.csdn.net/weixin_43064185/article/details/130820437
func TestC(t *testing.T) {
	var arr []*int
	for i := 0; i < 3; i++ {
		// arr = append(arr, &i) // bug

		j := i // fixed
		arr = append(arr, &j)
	}
	for _, f := range arr {
		fmt.Print(*f)
	}
}

// 不是bug
func TestC1(t *testing.T) {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i)
	}
}

// / bug
func TestC2(t *testing.T) {

	fns := make([]func(), 0)
	for i := 0; i < 3; i++ {
		// fns = append(fns, func() {  // bug
		// 	fmt.Print(i)
		// })

		j := i // fixed
		fns = append(fns, func() {
			fmt.Print(j)
		})
	}
	for _, f := range fns {
		f()
	}
}
