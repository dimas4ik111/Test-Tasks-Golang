// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
)

func binarySearch(num int, mas []int) bool {
	low := 0
	high := len(mas) - 1

	for low <= high {
		median := (low + high) / 2
		if mas[median] == num {
			fmt.Println(mas[low])
			return true
		}
		if mas[median] < num {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(mas) || mas[low] != num {
		return false
	}

	return true
}

func main() {
	mas := []int{512, 35, 243, 5432, 6423, 6, 2, 37, 547, 32, 786, 5, 83, 4, 89, 50654, 7285, 6, 5, 231, 54, 235, 436, 32, 6, 3246, 52, 7897, 52}
	fmt.Println("mas: \n", mas)
	var searchingValue int
	fmt.Scanln(&searchingValue)
	fmt.Println(binarySearch(searchingValue, mas))
}
