package main

import (
	"fmt"
)

func reverseString(input string) string {
	// Создаем пустую строку, в которую будем добавлять символы в обратном порядке
	reversed := ""

	// Преобразуем строку в массив рун
	runes := []rune(input)

	// Итерируемся по массиву рун с конца к началу
	for i := len(runes) - 1; i >= 0; i-- {
		// Добавляем руну в результирующую строку
		reversed += string(runes[i])
	}

	return reversed
}

func main() {
	input := "абырвалг"
	reversed := reverseString(input)
	fmt.Println("Исходная строка:", input)
	fmt.Println("Перевернутая строка:", reversed)
}
