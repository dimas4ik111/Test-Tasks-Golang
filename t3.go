package main

import (
	"fmt"
	"math"
	"sync"
)

func square(num int, wg *sync.WaitGroup, res *int) {
	*res += int(math.Pow(float64(num), 2))
	wg.Done()
}

func main() {
	mas := [5]int{2, 4, 6, 8, 10}
	res := 0

	var wg sync.WaitGroup

	for i := 0; i < len(mas); i++ {
		wg.Add(1)
		go square(mas[i], &wg, &res)
	}

	wg.Wait()

	fmt.Println("res = ", res)
}