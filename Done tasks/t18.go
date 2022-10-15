// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type counter struct {
	sync.RWMutex
	c uint64
}

func newCounter() *counter {
	return &counter{}
}

func (coun *counter) increment(wg *sync.WaitGroup) {
	coun.Lock()
	defer coun.Unlock()
	coun.c++
	wg.Done()
}

func main() {
	var (
		n uint64
		i uint64
		wg sync.WaitGroup
	)
	fmt.Scanln(&n)

	count := newCounter()

	for ; i < n; i++ {
		wg.Add(1)
		go count.increment(&wg)
	}
	wg.Wait()
	fmt.Println(count.c)
}
