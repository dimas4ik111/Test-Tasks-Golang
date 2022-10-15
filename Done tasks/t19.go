// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

package main

import (
	"fmt"
)

func reverseString(str *string) {
	runes := []rune(*str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*str = string(runes)
}

func main() {
	var str string
	fmt.Scanln(&str)
	reverseString(&str)
	fmt.Println(str)
}
