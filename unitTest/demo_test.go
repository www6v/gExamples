package main

import (
	"testing"
	//"strconv"
)

func TestFirstTry(t *testing.T) {
	t.Log("My first try!")
	t.Log(Monday, Tuesday)
	//var i = strconv.Itoa(1)
}

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

func TestFirstTry1(t *testing.T) {
	t.Log(Readable)
}
