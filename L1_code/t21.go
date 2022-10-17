// Реализовать паттерн «адаптер» на любом примере.

package main

import (
	"fmt"
	"strconv"
)

type Work interface {
	Addition()
}

func (n abInt) Addition() {
	fmt.Println(n.a + n.b)
}

type abInt struct {
	a int64
	b int64
}

type abStr struct {
	a string
	b string
}

type AdapterString struct {
	str *abStr
}

func (s *abStr) ConvertToInt() (int64, int64) {
	n1, _ := strconv.ParseInt(s.a, 10, 32)
	n2, _ := strconv.ParseInt(s.b, 10, 32)
	return n1, n2
}

func (adapt *AdapterString) Addition() {
	// a, b = adapt.str.ConvertToInt()
	a, b := adapt.str.ConvertToInt()
	fmt.Println(a + b)
}

func NewAdapter(a, b string) *AdapterString {
	return &AdapterString{
		str: newStr(a, b),
	}
}

func newInt(a, b int64) *abInt {
	return &abInt{
		a: a,
		b: b,
	}
}

func newStr(a, b string) *abStr {
	return &abStr{
		a: a,
		b: b,
	}
}

func main() {
	smtInt := newInt(100, 111)
	stradapt := NewAdapter("10", "11")
	fmt.Println("ABSInt")
	smtInt.Addition()
	fmt.Println("Adapter")
	stradapt.Addition()
}
