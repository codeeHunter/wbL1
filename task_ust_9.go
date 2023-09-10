package main

import "fmt"

func main() {
	// Slice
	// 1. Создание slice с помощью литерала
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	// 2. Создание slice с использованием функции make
	mySlice1 := make([]int, 3, 5) // Тип, длина и емкость
	fmt.Println(mySlice1)

	// 3. Создание slice из существующего массива или slice
	arr := []int{1, 2, 3, 4, 5}
	mySlice3 := arr[1:4] // От 2-го до 4-го элемента (не включая)
	fmt.Println(mySlice3)

	// Map
	// 1. Создание map с помощью литерала
	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(myMap)

	// 2. Создание map с использованием функции make
	myMap1 := make(map[string]int)
	myMap1["one"] = 1
	myMap1["two"] = 2
	myMap1["three"] = 3
	fmt.Println(myMap1)

}
