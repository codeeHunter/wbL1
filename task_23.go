package main

import (
	"fmt"
)

// Функция для удаления элемента из слайса по индексу
func removeElement(slice []int, indexToRemove int) []int {
	// Проверка на допустимость индекса
	if indexToRemove < 0 || indexToRemove >= len(slice) {
		return slice
	}

	// Удаление элемента
	copy(slice[indexToRemove:], slice[indexToRemove+1:])
	return slice[:len(slice)-1]
}

func main() {
	// Исходный слайс
	slice := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	indexToRemove := 2

	// Удаление элемента с помощью функции
	result := removeElement(slice, indexToRemove)

	// Вывод результата
	fmt.Println(result)
}
