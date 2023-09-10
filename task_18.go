package main

import (
	"fmt"
	"sync"
)

// Counter - структура для хранения счетчика и мьютекса для обеспечения безопасности доступа к нему.
type Counter struct {
	mu      sync.Mutex
	counter int
}

// Increment - инкрементирует счетчик на 1.
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

// Value - возвращает текущее значение счетчика.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func main() {
	counter := &Counter{}

	var wg sync.WaitGroup
	numRoutines := 100

	// Запускаем 100 горутин, которые инкрементируют счетчик.
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	// Ожидаем завершения всех горутин.
	wg.Wait()

	// Выводим итоговое значение счетчика.
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.Value())
}
