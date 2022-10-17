// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Worker(ch chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	ch := make(chan int)
	var (
		n                  int
		checkForExitSignal bool
	)
	_, err := fmt.Scanln(&n)
	checkForExitSignal = false
	if err != nil {
		fmt.Println("Error: Input value isn't type \"INT\" ")
		panic(err)
	}

	go func() {
		for {
			s := make(chan os.Signal)
			signal.Notify(s, syscall.SIGINT)
			switch <-s {
			case syscall.SIGINT:
				checkForExitSignal = true
			default:
				checkForExitSignal = false
			}
		}
	}()

	if n > 0 {
		for i := 0; i < n; i++ {
			go Worker(ch)
		}

		for {
			if !checkForExitSignal {
				ch <- rand.Int() % 10000
				} else {
					close(ch)
					time.Sleep(time.Millisecond * 10)
				fmt.Println("\nChanel close")
				break
			}
		}
	}
}
