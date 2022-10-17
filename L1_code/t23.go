package main

import "fmt"

func main() {
	mas := make([]int, 10)
	for i := 0; i < 10; i++ {
		mas[i] = i
	}
	fmt.Println(mas)
	var n int
	fmt.Print("index of delete element: ")
	fmt.Scanln(&n)
	if n > 9 {
		return
	}
	mas = append(mas[:n - 1], mas[n:]...)
	fmt.Println(mas)
}
