package main

import (
	"fmt"
)

func copyWriterFirstInSecond(fChan chan int, sChan chan int) {
	for num := range fChan {
		sChan <- num * 2
	}
}

func main() {
	var count int
	fmt.Scanln(&count)

	myMas := make([]int, count)

	for i := range myMas {
		fmt.Scanln(&myMas[i])
	}

	fChan := make(chan int)
	sChan := make(chan int)

	defer close(fChan)
	defer close(sChan)

	go copyWriterFirstInSecond(fChan, sChan)

	for i := range myMas {
		fChan <- myMas[i]
		fmt.Print(<-sChan, " ")
	}
}
