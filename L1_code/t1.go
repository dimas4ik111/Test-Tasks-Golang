// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age int
}

type Action struct {
	Human
}

func (h Human) viewInfo() string {
	// resStr := ""
	return "name: " + h.name + " age: " + strconv.Itoa(h.age)
}

func main() {
	h := Human{name: "Example", age: int(21)}
	fmt.Println(h.viewInfo())
}
