package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range dataCh {
		fmt.Printf("Воркер %d получил данные: %d\n", id, data)
	}
}

func main() {
	// Создаем канал для передачи данных от главного потока к воркерам
	dataCh := make(chan int)

	var wg sync.WaitGroup

	var N int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&N)

	// Запускаем воркеров
	for i := 1; i <= N; i++ {
		wg.Add(1)
		go worker(i, dataCh, &wg)
	}

	// Запускаем обработчик сигнала Ctrl+C для завершения программы
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM) // (Ctrl+C), Заверешнеия сигнала

	// Главный поток продолжает отправлять данные в канал
	go func() {
		for i := 1; i <= 10; i++ {
			dataCh <- i
			// Ждем некоторое время перед отправкой следующих данных
			time.Sleep(1 * time.Second)
		}

		// Ожидаем завершения работы всех воркеров
		wg.Wait()

		// Закрываем канал dataCh после завершения всех воркеров
		close(dataCh)
	}()

	// Ожидаем получения сигнала Ctrl+C
	<-sigCh

	fmt.Printf("Программа завершила работу.\n")
}
