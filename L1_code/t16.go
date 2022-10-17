// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func main() {
	m := []int{12, 52, 62, -1, 0, 66, 23, 555, -23, -77, 23}
	fmt.Println("before: ", m)

	sort.Ints(m)

	fmt.Println("after: ", m)
}
