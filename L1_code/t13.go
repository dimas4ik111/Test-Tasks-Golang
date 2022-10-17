// Поменять местами два числа без создания временной переменной.

package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b int
	)

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	defer fmt.Println("begin: a = ", a, "b = ", b)
	
	a, b = b, a
	
	fmt.Println("result: a = ", a, "b = ", b)
}
