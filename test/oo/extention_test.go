package oo

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

//type Dog struct {
//	p *Pet
//}
//
//func (d * Dog) Speak() {
//   fmt.Println("Wang")
//}
//
//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println(" ", host)
//}
//
//func TestDog(t *testing.T) {
//	dog:= new(Dog)
//	dog.SpeakTo("Chao")
//}

type Dog struct {
	Pet
}

func (p *Dog) Speak() {
	fmt.Print("Wang")
}

func TestDog(t *testing.T) {
	//var dog Pet := new(Dog) /// compile fail // 不支持显示类型转换， 不支持继承
	dog := new(Dog)
	dog.SpeakTo("Chao") /// 不支持重载， 不支持LSP
}
