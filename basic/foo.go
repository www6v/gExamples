package basic

import "fmt"

type Foo struct {
	Id   int
	Name string
}

func init() {
	f := &Foo{Id: 123, Name: "abc"}
	fmt.Printf("init foo object: %v\n", f)
}
