package main

import (
	"fmt"
	"testing"
)

type Bird interface {
	Fly()
	Fly1()
}
type Default struct {
}

func (d Default) Fly() {
	fmt.Println("default.fly")
}
func (d Default) Fly1() {
	fmt.Println("default.fly1")
}

type Animal struct {
}

type Free struct {
	Default
}

func (d *Free) Fly() {
	fmt.Println("Free.fly")
}

func (a Animal) Fly() {
	fmt.Println("animal.fly")
}
func (a Animal) Fly1() {
	fmt.Println("animal.fly1")
}
func TestExtends(t *testing.T) {
	var a Bird
	a = &Free{}
	a.Fly()
	a.Fly1()
}
