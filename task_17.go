package main

import (
	"fmt"
)

func main() {
	// Отсортированный массив
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7

	// Вызываем бинарный поиск
	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден по индексу %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}

// Функция для бинарного поиска
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid // Элемент найден, возвращаем индекс
		} else if arr[mid] < target {
			left = mid + 1 // Искать в правой половине
		} else {
			right = mid - 1 // Искать в левой половине
		}
	}

	return -1 // Элемент не найден
}
