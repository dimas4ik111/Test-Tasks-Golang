// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// var justString string

// func someFunc() {

//   v := createHugeString(1 << 10)

//   justString = v[:100]

// }

// func main() {

//   someFunc()

// }

package main

import "fmt"

func createHugeString(n int) string {
	var str string
	for i := 0; i < n; i++ {
		str += "😬"
		// str += "lol KEK "
	}
	return str
}

func someFunc(str string) []rune {

	str = createHugeString(1 << 10)

	// create []rune, чтобы не потерять символы которые весят более 1 байта
	justString := []rune(str)

	return justString[:100]
}

func main() {
	// Глобальная переменная может быть проблемой при многопоточности
	var justString string

	fmt.Println(string(someFunc(justString)))
}
