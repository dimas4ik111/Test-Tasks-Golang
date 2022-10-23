package main

import (
	"fmt"
	"time"
)

type Computer interface {
	TurnOn()
	TurnOf()
}

type Monitor struct {
}

type Keyboard struct {
}

type Mouse struct {
}

type PC struct {
	Monitor  *Monitor
	Keyboard *Keyboard
	Mouse    *Mouse
}

func (p *PC) TurnOn() {
	fmt.Println("Computer turn on...")
	time.Sleep(time.Duration(1) * time.Second)
	p.Monitor.TODO("Monitor on")
	p.Keyboard.TODO("Deyboard on")
	p.Mouse.TODO("Keyboard on")
}

func (p *PC) TurnOf() {
	fmt.Println("Computer turn of...")
	time.Sleep(time.Duration(1) * time.Second)
	p.Monitor.TODO("Monitor of")
	p.Keyboard.TODO("Deyboard of")
	p.Mouse.TODO("Keyboard of")
}

func (m Monitor) TODO(do string) {
	fmt.Println(do)
	time.Sleep(time.Duration(1) * time.Second)
}

func (k Keyboard) TODO(do string) {
	fmt.Println(do)
	time.Sleep(time.Duration(1) * time.Second)
}

func (m Mouse) TODO(do string) {
	fmt.Println(do)
	time.Sleep(time.Duration(1) * time.Second)
}

func NewPc() *PC {
	return &PC{
		Monitor:  &Monitor{},
		Keyboard: &Keyboard{},
		Mouse:    &Mouse{},
	}
}

func main() {
	comp := NewPc()
	comp.TurnOn()
	for i := 0; i < 10; i++ {
		fmt.Println("Work")
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
	comp.TurnOf()
}

// Паттерн Facade относится к структурным паттернам уровня объекта.

// Наша задача, сделать простой, единый интерфейс, через который можно было бы взаимодействовать с подсистемами.
