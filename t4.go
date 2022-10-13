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
