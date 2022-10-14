package main

import "fmt"

func main() {
	inputMas := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	myMap := make(map[int][]float64)

	for _, num := range inputMas {
		key := int(num) / 10 * 10
		myMap[key] = append(myMap[key], num)
	}

	fmt.Println(myMap)
}
