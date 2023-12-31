package main

import "fmt"

func main() {
	// Небуферизированные каналы предоставляют синхронный механизм передачи данных между горутинами.
	// Они ожидают, пока данные будут прочитаны из канала до того, как горутина-отправитель сможет продолжить выполнение.
	// Это гарантирует, что данные не будут потеряны и что обмен данными происходит точно в нужный момент.
	ch := make(chan int) // Создаем небуферизированный канал

	go func() {
		ch <- 42 // Передаем значение в канал
	}()

	value := <-ch      // Получаем значение из канала
	fmt.Println(value) // Выводим значение: 42

	// Буферизированные каналы позволяют более гибко управлять обменом данными между горутинами, так как они имеют фиксированную емкость.
	// Горутины-отправители могут отправлять данные в канал до того, как другие горутины начнут их получать.
	// Однако, если буфер полностью заполнен, отправитель будет заблокирован до тех пор, пока место не освободится.

	// Важно помнить, что при использовании буферизированных каналов нужно аккуратно следить за тем, чтобы не произошло "зависания" программы из-за блокировки горутин.
	ch1 := make(chan int, 2) // Создаем буферизированный канал с емкостью 2

	go func() {
		ch1 <- 1
		ch1 <- 2
	}()

	value1 := <-ch1
	value2 := <-ch1

	fmt.Println(value1, value2) // Выводим значения: 1 2
}
