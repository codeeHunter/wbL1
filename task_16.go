package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Создаем случайный массив чисел для сортировки
	rand.Seed(time.Now().UnixNano())
	arr := generateRandomArray(10)
	fmt.Println("Исходный массив:", arr)

	// Вызываем функцию для сортировки
	quicksort(arr)

	fmt.Println("Отсортированный массив:", arr)
}

// Функция для быстрой сортировки
func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	// Выбираем случайный опорный элемент
	pivotIndex := rand.Int() % len(arr)
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Разбиваем массив на две части
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Помещаем опорный элемент на правильное место
	arr[left], arr[right] = arr[right], arr[left]

	// Рекурсивно сортируем две части
	quicksort(arr[:left])
	quicksort(arr[left+1:])
}

// Функция для генерации случайного массива чисел
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}
