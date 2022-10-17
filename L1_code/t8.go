// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

func errMessage(ok error) {
	if ok != nil {
		panic("you write wrong format")
	}
}

func main() {
	fmt.Println("Hi")
	var number int64

	// number = 2
	fmt.Println("write your number(int64)")
	_, ok := fmt.Scanln(&number)
	errMessage(ok)

	var (
		iBit   int
		bitVal int
	)

	fmt.Println("n - bit:")
	_, ok = fmt.Scanln(&iBit)
	errMessage(ok)

	fmt.Println("value:")
	_, ok = fmt.Scanln(&bitVal)
	errMessage(ok)

	fmt.Println("Before num:", number)
	if bitVal == 1 {
		number |= 1 << iBit
		fmt.Println("After num:", number)
	} else if bitVal == 0 {
		number &^= 1 << iBit
		fmt.Println("After num:", number)
	} else {
		fmt.Println("ERROR, you must write 1 or 0")
	}
}
