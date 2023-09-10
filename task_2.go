package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем канал, в который будет отправляться результаты (квадраты чисел).
	results := make(chan int)

	// Создаем WaitGroup, чтобы дождаться завершения всех горутин.
	var wg sync.WaitGroup

	for _, num := range numbers {
		// Увеличиваем счетчик WaitGroup, чтобы указать, что мы запустили горутину.
		wg.Add(1)

		// Запускаем горутину, которая вычисляет квадрат числа.
		// Используется замыкание, чтобы передать значение числа (num) внутрь горутины.
		go func(x int) {
			defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины.
			square := x * x
			results <- square
		}(num)
	}

	// Запускаем еще одну горутину, которая будет ждать завершеня всех горутин.
	go func() {
		wg.Wait()
		close(results)
	}()

	for square := range results {
		fmt.Println(square)
	}
}
