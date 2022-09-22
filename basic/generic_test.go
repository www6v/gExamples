package basic

import (
	"fmt"
	"testing"
)

type C1 interface {
	// ~int | ~int32
	// int
	M1()
}

type T2 struct{}

func (T2) M1() {
}

type T1 int

func (T1) M1() {
}

func foo[P C1](t P)() {
	fmt.Print("abc")
}

/// gotip test generic_test.go
func TestGenericFuction(testing *testing.T) {
// func main() {	
	var t1 T1
	foo(t1)
	var t T2
	foo(t)
}
