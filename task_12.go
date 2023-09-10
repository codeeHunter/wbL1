package main

import "fmt"

func main() {
	// Создаем множество (map) для хранения строк
	set := make(map[string]bool)

	// Последовательность строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Добавляем каждую строку в множество
	for _, s := range sequence {
		set[s] = true
	}

	// Выводим уникальные строки
	fmt.Println("Уникальные строки в последовательности:")
	for key := range set {
		fmt.Println(key)
	}
}
