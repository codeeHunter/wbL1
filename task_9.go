package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем каналы для передачи данных между горутинами.
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	// Используем WaitGroup для ожидания завершения горутин.
	var wg sync.WaitGroup

	// Горутина для записи чисел в первый канал.
	go func() {
		defer close(inputChannel) // Закрываем канал после окончания записи.

		numbers := []int{1, 2, 3, 4, 5} // Здесь вы можете заменить на свой массив чисел.

		for _, num := range numbers {
			inputChannel <- num
		}
	}()

	// Горутина для умножения чисел и записи их во второй канал.
	go func() {
		defer close(outputChannel) // Закрываем канал после окончания записи.

		for x := range inputChannel {
			result := x * 2
			outputChannel <- result
		}

		// Уменьшаем счетчик ожидаемых горутин.
		wg.Done()
	}()

	// Горутина для вывода результатов в stdout.
	go func() {
		// Читаем данные из второго канала и выводим их в stdout.
		for result := range outputChannel {
			fmt.Println(result)
		}
	}()

	// Увеличиваем счетчик ожидаемых горутин на 1.
	wg.Add(1)

	// Ожидаем завершения горутины, которая выполняет умножение.
	wg.Wait()
}
