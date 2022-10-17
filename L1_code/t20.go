// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun - wons god nus».

package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(str *string) {
	runes := []rune(*str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*str = string(runes)
}

func reverseWordsIntStr(str *string) {
	runes := []rune(*str)
	masString := make([]string, 0)
	var buf string
	for _, v := range(runes) {
		if string(v) != " " {
			buf += string(v)
		}
		if string(v) == " " {
			masString = append(masString, buf)
			buf = ""
		}
	}
	masString = append(masString, buf)
	var result string
	for _, word := range(masString) {
		reverseString(&word)
		result += word + " "
	}
	*str = result
}

func main() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	s := scanner.Scan()
	if !s {
		return
	}

	str = scanner.Text()
	reverseWordsIntStr(&str)
	fmt.Println(str)
}
