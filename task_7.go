package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем мапу и мьютекс для обеспечения конкурентной записи
	data := make(map[string]int)
	var mu sync.Mutex

	// Канал для синхронизации горутин
	done := make(chan bool)

	// Количество горутин, которые будут записывать данные в мапу
	numGorutines := 5

	for i := 0; i < numGorutines; i++ {
		go func(id int) {
			key := fmt.Sprintf("Key%d", id)
			value := id * 10

			// Захватываем мьютекс для безопасной записи в мапу
			mu.Lock()
			data[key] = value
			mu.Unlock()

			fmt.Printf("Горутина %d записала в мапу: %s -> %d\n", id, key, value)

			// Отправляем сигнал о завершении горутины
			done <- true
		}(i)
	}

	for i := 0; i < numGorutines; i++ {
		<-done
	}

	// Выводим содержимое мапы после всех записей
	fmt.Println("Содержимое мапы:", data)

}
