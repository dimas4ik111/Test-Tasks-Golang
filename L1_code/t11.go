// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{8, 43, 12, 2, 5, 1, 3, 3, 3, 3, 4, 1, 456}

	helpMap := make(map[int]int)
	result := make([]int, 0)

	for _, num := range a {
		helpMap[num] = 1
	}

	for _, num := range b {
		_, ok := helpMap[num]
		if ok {
			helpMap[num] = 1
		} else {
			helpMap[num] += 1
		}
	}

	for key := range helpMap {
		result = append(result, key)
	}

	fmt.Println(result)
}
