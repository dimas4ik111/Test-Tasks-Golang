// Разработать программу, которая в рантайме 
// способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
)

func whatsType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println(t, "it's int")
	case float64:
		fmt.Println(t, "it's float64")
	case string:
		fmt.Println(t, "it's string")
	case bool:
		fmt.Println(t, "it's bool")
	case chan string:
		fmt.Println(t, "it's chan string")
	default:
		fmt.Println(t, "unknown type")
	}
}

func main() {
	ch := make(chan string)
	m := make(map[int]int)
	m[0] = 123
	defer close(ch)
	whatsType(int(1))
	whatsType(float64(55.44))
	whatsType("string")
	whatsType(false)
	whatsType(true)
	whatsType(ch)
	whatsType(m)
}
