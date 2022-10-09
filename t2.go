package main

import (
	"fmt"
	"math"
	"sync"
)

func square(num *int, wg *sync.WaitGroup) {
	*num = int(math.Pow(float64(*num), 2))
	wg.Done()
}

func main() {
	mas := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for i := 0; i < len(mas); i++ {
		wg.Add(1)
		go square(&mas[i], &wg)
	}

	wg.Wait()

	for i := 0; i < len(mas); i++ {
		if i != len(mas) - 1 {
			fmt.Print(mas[i])
			fmt.Print(" ")
		} else if i == len(mas) - 1 {
			fmt.Println(mas[i])
		}
	}
}
