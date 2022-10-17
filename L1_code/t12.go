// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество

package main

import (
	"fmt"
)

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println("begin:", str)

	myMap := make(map[string]int)
	result := make([]string, 0)

	for _, s := range(str) {
		_, ok := myMap[s]
		if ok {
			myMap[s]++
		} else {
			myMap[s] = 1
		}
		if myMap[s] == 1 {
			result = append(result, s)
		}
	}

	fmt.Println("result: ", result)
}
