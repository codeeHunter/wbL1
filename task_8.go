package main

import (
	"fmt"
)

func setBit(number int64, position uint, value uint) int64 {
	if value == 1 {
		// Устанавливаем бит в 1
		number |= 1 << position
	} else if value == 0 {
		// Устанавливаем бит в 0
		number &^= 1 << position
	} else {
		// Если value не 0 или 1, выводим сообщение об ошибке
		fmt.Println("Значение value должно быть 0 или 1")
	}
	return number
}

func main() {
	var number int64
	var position uint
	var value uint

	fmt.Print("Введите число (int64): ")
	_, _ = fmt.Scan(&number)
	fmt.Print("Введите позицию бита: ")
	_, _ = fmt.Scan(&position)
	fmt.Print("Введите значение (0 или 1): ")
	_, _ = fmt.Scan(&value)

	result := setBit(number, position, value)
	fmt.Printf("Результат: %d\n", result)
}
