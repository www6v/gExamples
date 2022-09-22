package basic

import "fmt"

const (
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
)

const (
	Readable = 1 << 2
	Writeable
	Execable
)

type Foo struct {
	Id   int
	Name string
}

func Test1() {
	f := &Foo{Id: 123, Name: "abc"}
	fmt.Printf("init foo object: %v\n", f)
}
