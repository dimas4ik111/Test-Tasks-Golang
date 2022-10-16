// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a1, b1 string
	fmt.Scanln(&a1)
	fmt.Scanln(&b1)
	result := new(big.Int)
	a, ok1 := new(big.Int).SetString(a1, 10)
	b, ok2 := new(big.Int).SetString(b1, 10)
	if !ok1 || !ok2 {
		panic("a, b, not numbers")
	}
	fmt.Println("a + b = ", result.Add(a, b))
	fmt.Println("a - b = ", result.Sub(a, b))
	fmt.Println("a / b = ", result.Div(a, b))
	fmt.Println("a * b = ", result.Mul(a, b))
}
