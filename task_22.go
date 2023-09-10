package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var a, b float64

	// Создаем новый сканер для чтения с консоли
	scanner := bufio.NewScanner(os.Stdin)

	// Ввод первого числа a
	fmt.Print("Введите первое число (a): ")
	if scanner.Scan() {
		input := scanner.Text()
		var err error
		a, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Ошибка ввода числа a:", err)
			return
		}
	}

	// Ввод второго числа b
	fmt.Print("Введите второе число (b): ")
	if scanner.Scan() {
		input := scanner.Text()
		var err error
		b, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Ошибка ввода числа b:", err)
			return
		}
	}

	if a > math.Pow(2, 20) && b > math.Pow(2, 20) {
		// Умножение
		multiplicationResult := a * b
		fmt.Printf("Умножение: %.2f * %.2f = %.2f\n", a, b, multiplicationResult)

		// Деление
		divisionResult := a / b
		fmt.Printf("Деление: %.2f / %.2f = %.2f\n", a, b, divisionResult)

		// Сложение
		additionResult := a + b
		fmt.Printf("Сложение: %.2f + %.2f = %.2f\n", a, b, additionResult)

		// Вычитание
		subtractionResult := a - b
		fmt.Printf("Вычитание: %.2f - %.2f = %.2f\n", a, b, subtractionResult)
	} else {
		fmt.Println("Оба числа должны быть больше или равны 2^20")
	}
}
