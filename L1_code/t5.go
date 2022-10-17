// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 1; ; i++ {
		ch <- i
	}
}

func reader(ch chan int) {
	for r := range ch {
		time.Sleep(time.Second - time.Millisecond)
		fmt.Println(r)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	
	ch := make(chan int)
	go sender(ch)
	go reader(ch)

	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
	}

	close(ch)
}
