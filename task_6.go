package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Метод 1: Используем канал для остановки горутины
	stopChan1 := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopChan1:
				fmt.Println("Горутина 1 остановлена посредством канала")
				return
			default:
				// Ваш код работы горутины здесь
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Метод 2: Используем контекст для остановки горутины
	ctx2, cancel2 := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("Горутина 2 остановлена посредством контекста")
				return
			default:
				// Ваш код работы горутины здесь
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Метод 3: Используем флаг для остановки горутины
	var shouldStop3 bool
	wg.Add(1)
	go func() {
		defer wg.Done()
		for !shouldStop3 {
			// Ваш код работы горутины здесь
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Горутина 3 остановлена посредством флага")
	}()

	// Создаем канал для остановки горутины
	stopChan2 := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина 4 начала выполнение")
		i := 0
		for {
			select {
			case <-stopChan2:
				fmt.Println("Горутина 4 остановлена по запросу")
				return
			default:
				i++
				fmt.Printf("Горутина 4: шаг %d\n", i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Ждем некоторое время перед остановкой горутины 4
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Отправляем сигнал на остановку горутины 4")
	close(stopChan2) // Отправляем сигнал остановки горутины 4

	// Завершение горутин 1, 2 и 3
	close(stopChan1)
	cancel2()
	shouldStop3 = true

	wg.Wait()
	fmt.Println("Все горутины завершились")
}
