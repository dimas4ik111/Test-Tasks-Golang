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

// Example of work with waitGroup
// sourse: https://gobyexample.com/waitgroups

// package main

// import (
//     "fmt"
//     "sync"
//     "time"
// )

// func worker(id int) {
//     fmt.Printf("Worker %d starting\n", id)

//     time.Sleep(time.Second)
//     fmt.Printf("Worker %d done\n", id)
// }

// func main() {

//     var wg sync.WaitGroup

//     for i := 1; i <= 5; i++ {
//         wg.Add(1)

//         i := i

//         go func() {
//             defer wg.Done()
//             worker(i)
//         }()
//     }

//     wg.Wait()

// }
